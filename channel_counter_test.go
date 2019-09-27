package main

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	c := NewCounter()
	c.IncrementCounter()
	c.DecrementCounter()
	fmt.Printf("\nCOUNT = %d\n", c.currentCount)
}

func BenchmarkChannel(b *testing.B) {
	c := NewCounter()
	for n := 0; n < 1000000; n++ {
		c.IncrementCounter()
		c.DecrementCounter()
	}
}
