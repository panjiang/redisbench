package models

import "time"

// SummaryResult : Summary result of all nodes
type SummaryResult struct {
	TotalTimes int64
	TotalDur   time.Duration
	TPS        int
}

// NodeResult : Only one node's result
type NodeResult struct {
	Order      int
	TotalTimes int64
	TotalDur   time.Duration
	TsBeg      time.Time
	TsEnd      time.Time
}
