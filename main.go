package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"NI/waste"
)

const Version = "1.5"

var (
	FlagCPU                    = flag.Duration("cpu", 0, "Interval of CPU streess to repeat. This flag is requried to emable CPU stress")
	FlagCPUduration            = flag.Duration("cpu-d", 2*time.Second, "Duration for each CPU stress cycle")
	FlagCPUpercent             = flag.Float64("cpu-p", 100.0, "Each CPU's load percentage to generate")
	FlagCPUcount               = flag.Int("cpu-n", runtime.NumCPU(), "Number of CPU cores to stress")
	FlagCPUglobalmaxpercent    = flag.Float64("cpu-m", 100.0, "Maximum limit of system's total CPU load percent to maintain. If other system processes consume CPU then this will reduce CPU utilization in attempt to maintain this limit.")
	FlagMemory                 = flag.Float64("mem", 0, "GiB of RAM memory to use")
	FlagMemoryRefresh          = flag.Bool("mem-r", false, "Enabling this refreshes the memeory every hour in attempt prevent OS from swap out the memoery and keep utilizing RAM.")
	FlagNetwork                = flag.Duration("net", 0, "Interval for network speed test")
	FlagNetworkConnectionCount = flag.Int("net-c", 10, "Set concurrent connections for network speed test")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Version", Version)
	fmt.Println("Platform:", runtime.GOOS, ",", runtime.GOARCH, ",", runtime.Version())

	flag.Parse()
	nothingEnabled := true

	if *FlagMemory != 0 {
		nothingEnabled = false
		fmt.Println("====================")
		fmt.Println("Starting memory consumption of", *FlagMemory, "GiB, with memory refresh set to", *FlagMemoryRefresh)
		go waste.Memory(*FlagMemory, *FlagMemoryRefresh)
		runtime.Gosched()
		fmt.Println("====================")
	}

	if *FlagCPU != 0 {
		nothingEnabled = false
		fmt.Println("====================")
		fmt.Println("Starting", *FlagCPUpercent, "% load of", *FlagCPUcount,"CPUs with interval of", *FlagCPU, ", min duration", *FlagCPUduration, "each, and max. system  load", *FlagCPUglobalmaxpercent,"%")
		go waste.CPU(*FlagCPU, *FlagCPUduration, *FlagCPUpercent, *FlagCPUcount, *FlagCPUglobalmaxpercent)
		runtime.Gosched()
		fmt.Println("====================")
	}

	if *FlagNetwork != 0 {
		nothingEnabled = false
		fmt.Println("====================")
		fmt.Println("Starting network speed testing with interval", *FlagNetwork)
		go waste.Network(*FlagNetwork, *FlagNetworkConnectionCount)
		runtime.Gosched()
		fmt.Println("====================")
	}

	if nothingEnabled {
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("\nFor duration or interval flags use time units. E.g., 1h5m10s, 5m, 1h, etc.")
	} else {
		// fatal error: all goroutines are asleep - deadlock!
		// select {} // fall asleep

		for {
			time.Sleep(24 * time.Hour)
		}
	}
}
