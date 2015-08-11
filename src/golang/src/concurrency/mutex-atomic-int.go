package main

import (
	"fmt"
	"sync"
)

// show A OMIT
type AtomicCounter struct { // HL
	mu sync.Mutex // only one goroutine can hold this. // HL
	n  int
}

// Add adds an integer to the counter.
func (c *AtomicCounter) Add(n int) {
	c.mu.Lock() // HL
	c.n += n
	c.mu.Unlock() // HL
}

func (c *AtomicCounter) Value() int {
	c.mu.Lock()         // HL
	defer c.mu.Unlock() // HL
	n := c.n            // What if the op here fails? defer helps. // HL
	return n
}

// end show A OMIT

func main() {
	var c AtomicCounter // HL

	done := make(chan struct{})
	go func() {
		c.Add(1) // one access // HL
		close(done)
	}()
	c.Add(1) // another concurrent access // HL
	<-done

	fmt.Println(c.Value()) // Definite output: 2 // HL
}
