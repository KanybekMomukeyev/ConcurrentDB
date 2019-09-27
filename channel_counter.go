package main

// ChannelCounter lalal
type ChannelCounter struct {
	currentCount    int
	updateChannel   chan func()
	responseChannel chan int
}

// NewCounter lalal
func NewCounter() *ChannelCounter {
	counter := new(ChannelCounter)
	counter.currentCount = 0
	counter.updateChannel = make(chan func(), 10000)
	counter.responseChannel = make(chan int, 10000)
	counter.waitForChannel()
	return counter
}

func (c *ChannelCounter) waitForChannel() {
	go func() {
		for f := range c.updateChannel {
			f()
		}
	}()
}

func (c *ChannelCounter) IncrementCounter() int {

	f := func() {
		c.currentCount = c.currentCount + 1
		c.responseChannel <- c.currentCount
	}

	c.updateChannel <- f

	select {
	case currentCount := <-c.responseChannel:
		return currentCount
	}
}

func (c *ChannelCounter) DecrementCounter() int {

	f := func() {
		c.currentCount = c.currentCount - 1
		c.responseChannel <- c.currentCount
	}

	c.updateChannel <- f

	select {
	case currentCount := <-c.responseChannel:
		return currentCount
	}
}
