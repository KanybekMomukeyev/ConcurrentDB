package main

import (
	"fmt"
	"testing"
)

func TestMutex(t *testing.T) {
	c := NewMutCounter()
	for i := 0; i < 10000; i++ {
		go func(index int, c *counter) {
			c.Add()
		}(i, c)
	}
	for i := 0; i < 1000; i++ {
		go func(index int, c *counter) {
			c.Minus()
		}(i, c)
	}
	fmt.Printf("\nCOUNT = %d\n", c.Get())
}

func BenchmarkMutex(b *testing.B) {
	c := NewMutCounter()
	for n := 0; n < 1000000; n++ {
		c.Add()
		c.Minus()
	}
}
