# CGroup Stats

`cgroup-stats` is a Go library designed for retrieving CPU and Memory quota information from Linux control groups (cgroups).

## Table of Contents

- [CGroup Stats](#cgroup-stats)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
    - [GetCPUQuota](#getcpuquota)
    - [GetMemoryLimit](#getmemorylimit)
  - [Example](#example)
  - [Contributing](#contributing)
  - [License](#license)

## Installation

To start using `cgroup-stats`, install Go and run `go get`:

```sh
go get -u github.com/mr-karan/cgroup-stats
```

## Usage

`cgroup-stats` provides easy-to-use functions to retrieve CPU and Memory quota information from cgroups.

### GetCPUQuota

`GetCPUQuota` function returns the CPU quota assigned to the cgroup.

```go
quota, err := cgroupstats.GetCPUQuota()
if err != nil {
	fmt.Println("Error getting CPU quota:", err)
	return
}
fmt.Println("CPU quota in the container:", quota)
```

### GetMemoryLimit

`GetMemoryLimit` function returns the memory limit assigned to the cgroup.

```go
memQuota, err := cgroupstats.GetMemoryLimit()
if err != nil {
	fmt.Println("Error getting Memory quota:", err)
	return
}
fmt.Println("Memory limit in the container:", memQuota/1048576)
```

## Example

Below is a simple example demonstrating the usage of `cgroup-stats`:

```go
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
		fmt.Println("Error getting Memory quota:", err)
		return
	}
	fmt.Println("Memory limit in the container:", memQuota/1048576)
}
```

## Contributing

We love contributions! Please review our [contribution guidelines](CONTRIBUTING.md) to get started.

## License

`cgroup-stats` is licensed under the MIT license. Please refer to the [LICENSE](LICENSE) file for detailed information.
