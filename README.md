# Redis & Redis Cluster redisbench Tool
* Can test redis single instance
* Can test redis cluster
* Supports running on multiple machines at the same time, for testing a large redis cluster

## Help
```
./redisbench -h
```

## Example

### Test single instance
```
./redisbench -a 192.168.10.11:6379 -c 500 -n 2000 -d 3
```

### Test cluster
```
./redisbench -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 500 -n 2000 -d 3
```

### Use multiple testing nodes
```
./redisbench -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 500 -n 2000 -d 3 -ma 192.168.10.11:9001,192.168.10.11:9002 -mo 1
./redisbench -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 500 -n 2000 -d 3 -ma 192.168.10.11:9001,192.168.10.11:9002 -mo 2
```
```
./redisbench -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 500 -n 2000 -d 3 -ma 192.168.10.11:9001,192.168.10.11:9002,192.168.10.11:9003 -mo 1
./redisbench -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 500 -n 2000 -d 3 -ma 192.168.10.11:9001,192.168.10.11:9002,192.168.10.11:9003 -mo 2
./redisbench -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 500 -n 2000 -d 3 -ma 192.168.10.11:9001,192.168.10.11:9002,192.168.10.11:9003 -mo 3
```

