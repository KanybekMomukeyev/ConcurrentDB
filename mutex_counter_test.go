package main

import (
	"fmt"
	"testing"
)

func TestMutex(t *testing.T) {
	c := NewMutCounter()
	c.Add()
	c.Add()
	c.Add()
	c.Add()
	c.Add()
	c.Add()
	c.Add()
	c.Add()
	c.Add()
	c.Add()
	c.Add()
	c.Add()

	c.Minus()
	c.Minus()
	c.Minus()
	c.Minus()
	c.Minus()
	c.Minus()
	c.Minus()

	fmt.Printf("\nCOUNT = %d\n", c.Get())
}

func BenchmarkMutex(b *testing.B) {
	c := NewMutCounter()
	for n := 0; n < 1000000; n++ {
		c.Add()
	}
	for n := 0; n < 100000; n++ {
		c.Minus()
	}
}
