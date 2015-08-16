// The following is a three stage pipeline.
package main

import (
	"fmt"
	"sync"
)

// gen emits the specified numbers on a channel that it returns.
func gen(nums ...int) <-chan int { // HL
	out := make(chan int)
	go func() {
		for _, n := range nums { // HL
			out <- n // HL
		} // HL
		close(out) // HL
	}()
	return out
}

// sq reads numbers from a channel and emits the square of each on a channel
// that it returns.
func sq(in <-chan int) <-chan int { // HL
	out := make(chan int)
	go func() {
		for n := range in { // HL
			out <- n * n // HL
		} // HL
		close(out) // HL
	}()
	return out
}

// merge fans in the values from multiple channels.
func merge(cs ...<-chan int) <-chan int { // HL
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c { // HL
			out <- n // HL
		} // HL
		wg.Done() // HL
	}
	wg.Add(len(cs)) // HL
	for _, c := range cs {
		go output(c) // HL
	}

	go func() {
		wg.Wait()  // HL
		close(out) // HL
	}()
	return out
}

func main() {
	// Explicit singular usage.
	c := gen(2, 3)
	out := sq(c)

	fmt.Println(<-out) // HL
	fmt.Println(<-out) // HL

	// Spicier usage draining the channel.
	for n := range sq(sq(gen(1, 2, 3, 4, 5, 6))) { // HL
		fmt.Println(n)
	}

	// Fanned-in usage.
	in := gen(1, 2, 3, 4)
	a := sq(in)
	b := sq(in)

	for n := range merge(a, b) { // HL
		fmt.Println(n)
	}
}
