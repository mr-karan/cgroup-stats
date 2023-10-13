package main

import (
	"fmt"
	"runtime"

	cgroupstats "github.com/mr-karan/cgroup-stats"
)

func main() {
	displayHostCPUInfo()
	displayContainerCPUQuota()
	displayContainerMemoryLimit()
}

func displayHostCPUInfo() {
	fmt.Printf("Number of CPUs on host: %d\n", runtime.NumCPU())
	fmt.Printf("Number of operating system threads: %d\n", runtime.GOMAXPROCS(0))
}

func displayContainerCPUQuota() {
	quota, err := cgroupstats.GetCPULimits()
	if handleError("CPU quota", err) {
		return
	}
	if quota == -1 {
		fmt.Printf("CPU quota is not set\n")
		return
	}
	fmt.Printf("CPU quota in the container: %f\n", quota)
}

func displayContainerMemoryLimit() {
	memQuota, err := cgroupstats.GetMemoryLimits()
	if handleError("Memory quota", err) {
		return
	}
	if memQuota == -1 {
		fmt.Printf("Memory limit is not set\n")
		return
	}
	fmt.Printf("Memory limit in the container: %f MiB\n", memQuota/1048576)
}

func handleError(resource string, err error) bool {
	if err != nil {
		fmt.Printf("Error getting %s: %s\n", resource, err)
		return true
	}
	return false
}
