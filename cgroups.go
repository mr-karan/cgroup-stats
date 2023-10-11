package cgroupstats

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// extractCPUQuota reads and parses the CPU quota from the cpu.max file.
func extractCPUQuota() (float64, error) {
	file, err := os.Open(cpuMaxFile)
	if err != nil {
		return 0, fmt.Errorf("failed to open %s: %w", cpuMaxFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() { // Only read the first line
		fields := strings.Fields(scanner.Text())
		if len(fields) < 1 || len(fields) > 2 {
			return 0, errors.New("invalid format in cpu.max")
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
				return 0, errors.New("invalid period in cpu.max")
			}
		}

		return float64(quota) / float64(period), nil
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scanner error: %w", err)
	}

	return -1, errors.New("cpu.max file is empty")
}

// extractMemQuota reads and parses the memory max limits from the memory.max file.
func extractMemQuota() (float64, error) {
	file, err := os.Open(memMaxFile)
	if err != nil {
		return 0, fmt.Errorf("failed to open %s: %w", memMaxFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() { // Only read the first line
		fields := strings.Fields(scanner.Text())
		if len(fields) != 1 {
			return 0, errors.New("invalid format in memory.max")
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

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scanner error: %w", err)
	}

	return -1, errors.New("memory.max file is empty")
}

// detectCgroupsV2 checks whether cgroups v2 is used by scanning the mountinfo file.
func detectCgroupsV2() (bool, error) {
	file, err := os.Open(procFileMount)
	if err != nil {
		return false, fmt.Errorf("failed to open %s: %w", procFileMount, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), cgroupV2Indicator) {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, fmt.Errorf("scanner error: %w", err)
	}

	return false, nil // Default to cgroups v1 if v2 is not found
}
