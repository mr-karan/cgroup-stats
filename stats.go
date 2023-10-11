package cgroupstats

import (
	"errors"
	"fmt"
)

const (
	procFileMount     = "/proc/self/mountinfo"
	cpuPeriodDefault  = 100000 // 100ms default
	cpuMaxFile        = "/sys/fs/cgroup/cpu.max"
	memMaxFile        = "/sys/fs/cgroup/memory.max"
	cgroupV2Indicator = "cgroup2"
)

// GetCPUQuota retrieves the CPU quota for cgroup v2.
// Returns -1 if no quota is set, or if cgroup v1 is detected.
func GetCPUQuota() (float64, error) {
	v2, err := detectCgroupsV2()
	if err != nil {
		return -1, fmt.Errorf("error checking cgroups version: %w", err)
	}

	if !v2 {
		return -1, errors.New("cgroups v1 is not supported yet")
	}

	return extractCPUQuota()
}

// GetMemoryLimit retrieves the max memory limits for cgroup v2.
// Returns -1 if no limit is set, or if cgroup v1 is detected.
func GetMemoryLimit() (float64, error) {
	v2, err := detectCgroupsV2()
	if err != nil {
		return -1, fmt.Errorf("error checking cgroups version: %w", err)
	}

	if !v2 {
		return -1, errors.New("cgroups v1 is not supported yet")
	}

	return extractMemQuota()
}
