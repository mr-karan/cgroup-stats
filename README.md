# CGroup Stats

`cgroup-stats` is a Go library designed for retrieving CPU and Memory limits information from Linux control groups (cgroups).

**NOTE**: cgroups v1 is not supported at the moment.

## Table of Contents

- [CGroup Stats](#cgroup-stats)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Demo](#demo)
    - [Building the Docker Image](#building-the-docker-image)
    - [Running Without Limits](#running-without-limits)
    - [Running With Limits](#running-with-limits)
  - [Usage](#usage)
    - [GetCPULimits](#getcpulimits)
    - [GetMemoryLimits](#getmemorylimits)
  - [Example](#example)

## Installation

To start using `cgroup-stats`, install Go and run `go get`:

```sh
go get -u github.com/mr-karan/cgroup-stats
```


## Demo

### Building the Docker Image

First, build the Docker image using the following command:

```bash
docker buildx build -t mrkaran/cgroups -f Dockerfile .
```

### Running Without Limits

Execute the container without any specific CPU or memory limits:

```bash
docker run --rm -it mrkaran/cgroups
```

You should see an output similar to:

```
Number of CPUs on host: 10
Number of operating system threads: 10
CPU quota is not set
Memory limit is not set
```

### Running With Limits

Now, let's run the container with specific CPU and memory limits. In this case, we're limiting the container to use only half a CPU core (`--cpus=0.5`) and 1000 MiB of memory (`--memory=1000m`):

```bash
docker run --rm -it --cpus=0.5 --memory=1000m mrkaran/cgroups
```

You should see an output reflecting the imposed limits:

```
Number of CPUs on host: 10
Number of operating system threads: 10
CPU quota in the container: 0.500000
Memory limit in the container: 1000.000000 MiB
```


## Usage

`cgroup-stats` provides easy-to-use functions to retrieve CPU and Memory quota information from cgroups.

### GetCPULimits

`GetCPULimits` function returns the CPU quota assigned to the cgroup.

```go
quota, _ := cgroupstats.GetCPULimits()
fmt.Println("CPU quota in the container:", quota)
```

### GetMemoryLimits

`GetMemoryLimits` function returns the memory limit assigned to the cgroup.

```go
memQuota, _ := cgroupstats.GetMemoryLimits()
fmt.Println("Memory limit in the container:", memQuota/1048576)
```

## Example

See [examples/main.go](examples/main.go) for a complete example.
