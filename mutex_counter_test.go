package main

import (
	"fmt"
	"testing"
)

func TestMutex(t *testing.T) {
	c := NewMutCounter()
	c.Add()
	c.Minus()
	fmt.Printf("\nCOUNT = %d\n", c.Get())
}

func BenchmarkMutex(b *testing.B) {
	c := NewMutCounter()
	for n := 0; n < 1000000; n++ {
		c.Add()
		c.Minus()
	}
}
