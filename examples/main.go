package main

import (
	"fmt"
	"runtime"

	cgroupstats "github.com/mr-karan/cgroup-stats"
)

func main() {
	numCPU := runtime.NumCPU()
	fmt.Println("Number of CPU on host:", numCPU)
	fmt.Println("number of operating system threads:", runtime.GOMAXPROCS(0))
	quota, err := cgroupstats.GetCPUQuota()
	if err != nil {
		fmt.Println("Error getting CPU quota:", err)
		return
	}
	fmt.Println("CPU quota in the container:", quota)

	memQuota, err := cgroupstats.GetMemoryLimit()
	if err != nil {
		fmt.Println("Error getting CPU quota:", err)
		return
	}
	fmt.Println("Memory limit in the container:", memQuota/1048576)
}
