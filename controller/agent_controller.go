package Controller

import (
	"context"
	"fmt"
	"github.com/coopersec/infomation-agent/models"
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/yahoo/vssh"
)

var (
	lastNetIOStatTimeStamp int64
	lastNetInfo            *models.NetInfo
	collectSysInfoTopic    string
)

func getCpuInfo() {
	var cpuInfo = new(models.CpuInfo)
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("cpu percent:%v\n", percent)
	cpuInfo.CpuPercent = percent[0]
}

func getMemInfo() {
	var memInfo = new(models.MemInfo)
	info, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("get mem info failed, err:%v", err)
		return
	}
	memInfo.Total = info.Total
	memInfo.Available = info.Available
	memInfo.Used = info.Used
	memInfo.UsedPercent = info.UsedPercent
	memInfo.Buffers = info.Buffers
	memInfo.Cached = info.Cached
}

func getDiskInfo() {
	var diskInfo = &models.DiskInfo{
		PartitionUsageStat: make(map[string]*models.UsageStat, 16),
	}
	parts, _ := disk.Partitions(true)
	for _, part := range parts {
		usageStatInfo, err := disk.Usage(part.Mountpoint) // 传挂载点
		if err != nil {
			fmt.Printf("get %s usage stat failed, err:%v", err)
			continue
		}
		usageStat := &models.UsageStat{
			Path:              usageStatInfo.Path,
			Fstype:            usageStatInfo.Fstype,
			Total:             usageStatInfo.Total,
			Free:              usageStatInfo.Free,
			Used:              usageStatInfo.Used,
			UsedPercent:       usageStatInfo.UsedPercent,
			InodesTotal:       usageStatInfo.InodesTotal,
			InodesUsed:        usageStatInfo.InodesUsed,
			InodesFree:        usageStatInfo.InodesFree,
			InodesUsedPercent: usageStatInfo.InodesUsedPercent,
		}
		diskInfo.PartitionUsageStat[part.Mountpoint] = usageStat
	}
}

func getNetInfo() {
	var netInfo = &models.NetInfo{
		NetIOCountersStat: make(map[string]*models.IOStat, 8),
	}
	currentTimeStamp := time.Now().Unix()
	netIOs, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("get net io counters failed, err:%v", err)
		return
	}
	for _, netIO := range netIOs {
		var ioStat = new(models.IOStat)
		ioStat.BytesSent = netIO.BytesSent
		ioStat.BytesRecv = netIO.BytesRecv
		ioStat.PacketsSent = netIO.PacketsSent
		ioStat.PacketsRecv = netIO.PacketsRecv
		netInfo.NetIOCountersStat[netIO.Name] = ioStat

		if lastNetIOStatTimeStamp == 0 || lastNetInfo == nil {
			continue
		}
		interval := currentTimeStamp - lastNetIOStatTimeStamp

		ioStat.BytesSentRate = (float64(ioStat.BytesSent) - float64(lastNetInfo.NetIOCountersStat[netIO.Name].BytesSent)) / float64(interval)

		ioStat.BytesRecvRate = (float64(ioStat.BytesRecv) - float64(lastNetInfo.NetIOCountersStat[netIO.Name].BytesRecv)) / float64(interval)
		ioStat.PacketsSentRate = (float64(ioStat.PacketsSent) - float64(lastNetInfo.NetIOCountersStat[netIO.Name].PacketsSent)) / float64(interval)
		ioStat.PacketsRecvRate = (float64(ioStat.PacketsRecv) - float64(lastNetInfo.NetIOCountersStat[netIO.Name].PacketsRecv)) / float64(interval)

	}
	lastNetIOStatTimeStamp = currentTimeStamp
	lastNetInfo = netInfo
}

// TODO: ERROR HANDLING, RETURN
func ExecuteCMD(private_key string, cmdStr string, ip string) (string, error) {
	vs := vssh.New().Start()
	config, _ := vssh.GetConfigPEM("vssh", private_key)
	vs.AddClient(ip, config, vssh.SetMaxSessions(4))
	vs.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := cmdStr
	timeout, _ := time.ParseDuration("6s")
	respChan := vs.Run(ctx, cmd, timeout)

	resp := <-respChan
	if err := resp.Err(); err != nil {
		log.Fatal(err)
	}

	stream := resp.GetStream()
	defer stream.Close()

	for stream.ScanStdout() {
		txt := stream.TextStdout()
		fmt.Println(txt)
	}
	return "", nil
}
