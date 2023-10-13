package cgroupstats

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	procFileMount     = "/proc/self/mountinfo"
	cpuPeriodDefault  = 100000 // 100ms default
	cpuMaxFile        = "/sys/fs/cgroup/cpu.max"
	memMaxFile        = "/sys/fs/cgroup/memory.max"
	cgroupV2Indicator = "cgroup2"
)

// fetchLine reads the first line from a file.
func fetchLine(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open %s: %w", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scanner error: %w", err)
	}
	return "", fmt.Errorf("no content in %s", filePath)
}

func extractCPUQuota() (float64, error) {
	content, err := fetchLine(cpuMaxFile)
	if err != nil {
		return 0, err
	}

	fields := strings.Fields(content)
	if len(fields) < 1 || len(fields) > 2 {
		return 0, fmt.Errorf("invalid format in cpu.max")
	}

	if fields[0] == "max" {
		return -1, nil // No quota is set
	}

	quota, err := strconv.Atoi(fields[0])
	if err != nil {
		return 0, fmt.Errorf("failed to convert quota to int: %w", err)
	}

	period := cpuPeriodDefault // Set default period
	if len(fields) == 2 {      // If period is provided, use it
		period, err = strconv.Atoi(fields[1])
		if err != nil {
			return 0, fmt.Errorf("failed to convert period to int: %w", err)
		}
		if period == 0 {
			return 0, fmt.Errorf("invalid period in cpu.max")
		}
	}

	return float64(quota) / float64(period), nil
}

func extractMemQuota() (float64, error) {
	content, err := fetchLine(memMaxFile)
	if err != nil {
		return 0, err
	}

	fields := strings.Fields(content)
	if len(fields) != 1 {
		return 0, fmt.Errorf("invalid format in memory.max")
	}

	if fields[0] == "max" {
		return -1, nil // No quota is set
	}

	quota, err := strconv.Atoi(fields[0])
	if err != nil {
		return 0, fmt.Errorf("failed to convert limit to int: %w", err)
	}

	return float64(quota), nil
}

// ensureCgroupV2 checks if cgroups v2 is being used by scanning the mountinfo file.
// It returns an error if cgroups v1 is detected or if there's an issue checking the version.
func ensureCgroupV2() error {
	file, err := os.Open(procFileMount)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", procFileMount, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), cgroupV2Indicator) {
			return nil // cgroups v2 detected, return without error
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	return fmt.Errorf("cgroups v1 is not supported yet") // Default to error if cgroups v2 is not detected
}
