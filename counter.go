package main

import "sync/atomic"

var counter uint64 = 0

func GetCounter() uint64 {
	return atomic.AddUint64(&counter, 1)
}
