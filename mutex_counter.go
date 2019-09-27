package main

import "sync"

func NewMutCounter() *counter {
	counter := new(counter)
	counter.n = 0
	counter.mu = sync.Mutex{}
	return counter
}

type counter struct {
	mu sync.Mutex
	n  int
}

func (c *counter) Add() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

func (c *counter) Minus() {
	c.mu.Lock()
	c.n--
	c.mu.Unlock()
}

func (c *counter) Get() int {
	c.mu.Lock()
	n := c.n
	c.mu.Unlock()
	return n
}

func (c *counter) Reset() {
	c.mu.Lock()
	if c.n > 8190 {
		c.n = 0
	}
	c.mu.Unlock()
}
