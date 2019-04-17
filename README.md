# Redis & Redis Cluster benchmark Tool

- Written in Golang
- Can test redis single instance
- Can test redis cluster
- Can take advantage of multi-core
- Supports running on multiple machines at the same time, for testing a large redis cluster (The same hardware of machines are needed for )

## Warning

Testing data keys named like `benchmark.set.*`, make sure they are not conflicting with your keys.

## Help

```console
$ ./redisbench -h

  -a string
        Redis instance address or Cluster addresses. IP:PORT[,IP:PORT]
  -c int
        Clients number for concurrence (default 1)
  -cluster
        true: cluster mode, false: instance mode
  -d int
        Data size in bytes (default 1000)
  -db int
        Choose a db, only for non-cluster (default 0)
  -ma string
        addresses for run multiple testers at the same time
  -mo int
        the order current tester is in multiple testers
  -n int
        Testing times at every client (default 3000)
  -p string
        The password for auth, only for non-cluster
```

## Example

Make sure your are testing an unused Redis, Because the tool will write lots of testing data into it.

### Test single instance

```console
$ ./redisbench -a 127.0.0.1:6379 -c 10 -n 2000 -d 1000


2019/03/01 14:24:34 Go...
2019/03/01 14:24:34 # BENCHMARK SINGLE (localhost:6379, db:0)
2019/03/01 14:24:34 * Clients Number: 10, Testing Times: 2000, Data Size(B): 1000
2019/03/01 14:24:34 * Total Times: 20000, Total Size(B): 20000000
2019/03/01 14:24:36 # BENCHMARK DONE
2019/03/01 14:24:36 * TIMES: 20000, DUR(s): 1.762, TPS(Hz): 11350

```

Redis keys like `benchmark.set.{client_id}.{test_times}`:

```
...
19996) "benchmark.set.3.9"
19997) "benchmark.set.2.1394"
19998) "benchmark.set.0.846"
19999) "benchmark.set.3.1690"
20000) "benchmark.set.8.848"
```

### Test cluster

```console
$ ./redisbench -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 10 -n 2000 -d 1000
```

### Use multiple testing nodes

```console
$ ./redisbench -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 10 -n 2000 -d 1000 -ma 192.168.10.11:9001,192.168.10.11:9002 -mo 1
$ ./redisbench -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 10 -n 2000 -d 1000 -ma 192.168.10.11:9001,192.168.10.11:9002 -mo 2
```

```console
$ ./redisbench -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 10 -n 2000 -d 1000 -ma 192.168.10.11:9001,192.168.10.11:9002,192.168.10.11:9003 -mo 1
$ ./redisbench -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 10 -n 2000 -d 1000 -ma 192.168.10.11:9001,192.168.10.11:9002,192.168.10.11:9003 -mo 2
$ ./redisbench -cluster=true -a 192.168.10.11:7000,192.168.10.11:7001 -c 10 -n 2000 -d 1000 -ma 192.168.10.11:9001,192.168.10.11:9002,192.168.10.11:9003 -mo 3
```
