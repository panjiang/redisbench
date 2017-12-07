# Redis & Redis Cluster Benchmark Tool

## Help
```
benchmark.go -h
```

## Example

### Test single instance
```
./benchmark -a 192.168.10.11:6379 -c 500 -n 2000 -d 3
```

### Test cluster
```
./benchmark -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 500 -n 2000 -d 3
```

### Use multiple testing nodes
```
./benchmark -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 500 -n 2000 -d 3 -ma 192.168.10.11:9001,192.168.10.11:9002 -mo 1
./benchmark -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 500 -n 2000 -d 3 -ma 192.168.10.11:9001,192.168.10.11:9002 -mo 2
```
```
./benchmark -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 500 -n 2000 -d 3 -ma 192.168.10.11:9001,192.168.10.11:9002,192.168.10.11:9003 -mo 1
./benchmark -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 500 -n 2000 -d 3 -ma 192.168.10.11:9001,192.168.10.11:9002,192.168.10.11:9003 -mo 2
./benchmark -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 500 -n 2000 -d 3 -ma 192.168.10.11:9001,192.168.10.11:9002,192.168.10.11:9003 -mo 3
```

