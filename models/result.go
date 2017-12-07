package models

// SummaryResult : Summary result of all nodes
type SummaryResult struct {
	TotalTimes int64
	TotalDur   int64
	TPS        int
}

// NodeResult : Only one node's result
type NodeResult struct {
	Order      int
	TotalTimes int64
	TotalDur   int64
	TsBeg      int64
	TsEnd      int64
}
