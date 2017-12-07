package tester

import "sync"

// Wg : Global wait group
var Wg sync.WaitGroup

// Multi : multiple tester contronl handler
var Multi *MultiTester
