package cgroupstats

// GetCPULimits retrieves the CPU quota for the cgroup.
// Returns -1 if no quota is set, or if an error occurs.
func GetCPULimits() (float64, error) {
	if err := ensureCgroupV2(); err != nil {
		return -1, err
	}
	return extractCPUQuota()
}

// GetMemoryLimits retrieves the max memory limits for cgroup v2.
// Returns -1 if no limit is set, or if an error occurs.
func GetMemoryLimits() (float64, error) {
	if err := ensureCgroupV2(); err != nil {
		return -1, err
	}
	return extractMemQuota()
}
