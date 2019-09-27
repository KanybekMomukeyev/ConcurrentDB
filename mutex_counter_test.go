package main

import (
	"fmt"
	"testing"
)

var (
	mutexCounter = NewMutCounter()
)

func TestMutex(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go func(index int, mutexCounter *counter) {
			mutexCounter.Add()
		}(i, mutexCounter)
	}
	for i := 0; i < 1000; i++ {
		go func(index int, mutexCounter *counter) {
			mutexCounter.Minus()
		}(i, mutexCounter)
	}
	fmt.Printf("\nCOUNT = %d\n", mutexCounter.Get())
}

func BenchmarkMutex(b *testing.B) {
	c := NewMutCounter()
	for n := 0; n < 1000000; n++ {
		c.Add()
		c.Minus()
	}
}
