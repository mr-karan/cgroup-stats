package cgroupstats

import "errors"

// Generic function to extract a quota
func extractQuota[T comparable](filePath string) (T, error) {
  // Implementation...
}

// GetQuota retrieves the quota for the given resource
func GetQuota[T comparable](resource string) (T, error) {
  v2, err := detectCgroupsV2()
  if err != nil {
    return *new(T), fmt.Errorf("error checking cgroups version: %w", err) 
  }

  if !v2 {
    return *new(T), errors.New("cgroups v1 not supported")
  }

  var filePath string
  switch resource {
  case "cpu":
    filePath = cpuMaxFile
  case "memory":
    filePath = memMaxFile
  default:
    return *new(T), fmt.Errorf("unknown resource: %s", resource)
  }

  return extractQuota[T](filePath)
}

// Usage:

cpuQuota, err := GetQuota[float64]("cpu")
if err != nil {
  // Handle error
}

memQuota, err := GetQuota[int64]("memory") 
if err != nil {
  // Handle error
}
