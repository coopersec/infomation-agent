package main

import (
	"fmt"
	"github.com/p3tr0v/chacal/antidebug"
	"github.com/p3tr0v/chacal/antimem"
	"github.com/p3tr0v/chacal/antivm"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	// VM Check
	if antidebug.ByProcessWatcher() { // Whether some debugger program founded, enter here.
		// exit or wait
		os.Exit(3)
	}
	if antimem.ByMemWatcher() { // Whether some program used for inspect memory founded, enter here.
		// exit or wait
		os.Exit(3)
	}
	if antivm.BySizeDisk(100) { // whether total disk size is less than 100 GB, enter here. You chose the size, always in GB.
		// exit or wait
		os.Exit(3)
	}
	if antivm.IsVirtualDisk() { // If Chacal guess you are on virtual disk, enter here.
		// exit or wait
		os.Exit(3)
	}
	if antivm.ByMacAddress() { // If Chacal guess you are on virtual MAC Address, enter here.
		// exit or wait
		os.Exit(3)
	}

	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get("localhost:3000/test")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)

}
