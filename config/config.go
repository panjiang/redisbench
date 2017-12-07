package config

import "flag"

// RedisAddr : Redis address or Cluster addresses
var RedisAddr string

// ClientNum : Client number for concurrence
var ClientNum int

// TestTimes : Test times of every client
var TestTimes int

// DataSize : Set data size at once
var DataSize int

// ClusterMode : Whether it is cluster mode
var ClusterMode bool

// MultiAddr : Run multi testers at the same time
// while single machine can not hold the testing
var MultiAddr string

// MultiOrder : The order current tester is
var MultiOrder int

// Parse configure from command line flags
func Parse() {
	flag.StringVar(&RedisAddr, "a", "", "Redis instance address or Cluster addresses. IP:PORT[,IP:PORT]")
	flag.IntVar(&ClientNum, "c", 1, "Clients number for concurrence")
	flag.IntVar(&TestTimes, "n", 1, "Testing times at every client")
	flag.IntVar(&DataSize, "d", 1000, "Data size in bytes")
	flag.BoolVar(&ClusterMode, "cluster", false, "true: cluster mode, false: instance mode")
	flag.StringVar(&MultiAddr, "ma", "", "addresses for run multiple testers at the same time")
	flag.IntVar(&MultiOrder, "mo", 0, "the order current tester is in multiple testers")
	flag.Parse()
}
