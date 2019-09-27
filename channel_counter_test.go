package main

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	c := NewCounter()
	for i := 0; i < 10000; i++ {
		go func(index int, c *ChannelCounter) {
			c.IncrementCounter()
		}(i, c)
	}
	for i := 0; i < 1000; i++ {
		go func(index int, c *ChannelCounter) {
			c.DecrementCounter()
		}(i, c)
	}
	fmt.Printf("\nCOUNT = %d\n", c.currentCount)
}

func BenchmarkChannel(b *testing.B) {
	c := NewCounter()
	for n := 0; n < 1000000; n++ {
		c.IncrementCounter()
		c.DecrementCounter()
	}
}
