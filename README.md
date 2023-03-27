# NI system strees utility

Cross-platform system stress utility to generate controlled system load for CPU, Memory and Network bandwidth. 
The systhentic load can be controlled by time intervals, load percents and amounts of resources to use.

Available for Windows, Linux and MacOS for both x86-64, ARM64 architectures.

## Usage


```shell
./NI -cpu 2h5m30s -cpu-d 15m -cpu-p 50 -cpu-m 80 -mem 2.5 -net 4h
```

```
  -cpu duration
        Interval of CPU streess (enables CPU stress)
  -cpu-d duration
        Min. duration for each CPU stress (default 2s)
  -cpu-m float
        Max limit of system's total CPU load perceent (default 100)
  -cpu-n int
        Number of CPU cores to stress (default AllCores)
  -cpu-p float
        Each CPU's load percentage (default 100)
  -mem float
        GiB of memory to use
  -net duration
        Interval for network speed test
  -net-c int
        Set concurrent connections for network speed test (default 10)
```
