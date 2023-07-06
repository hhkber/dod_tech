package main

import (
	"testing"
)

func TestDataRace(t *testing.T) {
	var i int64
	go func() { i++ }()
	i++
	//go func() { atomic.AddInt64(&i, 1) }()
	//atomic.AddInt64(&i, 1)
}
