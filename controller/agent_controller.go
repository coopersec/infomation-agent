package Controller

import (
	"fmt"
	"github.com/coopersec/infomation-agent/models"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
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

func GetMemInfo() uint64 {
	var memInfo = new(models.MemInfo)
	info, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("get mem info failed, err:%v", err)
		return 12
	}
	memInfo.Total = info.Total
	memInfo.Available = info.Available
	memInfo.Used = info.Used
	memInfo.UsedPercent = info.UsedPercent
	memInfo.Buffers = info.Buffers
	memInfo.Cached = info.Cached
	fmt.Println((memInfo.Total / 1000000000), "GB")
	return memInfo.Total
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
func ExecuteCMD() {
	PS, err := exec.LookPath("dir")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(PS)
	command := []string{"dir"}
	env := os.Environ()
	err = syscall.Exec("CMD", command, env)
	fmt.Println(err)
}

func Test()
