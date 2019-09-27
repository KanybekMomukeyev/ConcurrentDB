package main

import (
	"fmt"
	"testing"
)

var (
	channelCounter = NewCounter()
)

func TestChannel(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go func(index int, channelCounter *ChannelCounter) {
			channelCounter.IncrementCounter()
		}(i, channelCounter)
	}
	for i := 0; i < 1000; i++ {
		go func(index int, channelCounter *ChannelCounter) {
			channelCounter.DecrementCounter()
		}(i, channelCounter)
	}
	fmt.Printf("\nCOUNT = %d\n", channelCounter.GetCurrentCount())
}

func BenchmarkChannel(b *testing.B) {
	c := NewCounter()
	for n := 0; n < 1000000; n++ {
		c.IncrementCounter()
		c.DecrementCounter()
	}
}
