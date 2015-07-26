// Concurrency vs parallelism.
//
// An example demonstrating how Go <1.5 schedules goroutines on a single
// process, which means only one goroutine is running at any given time
// by default. The scheduler, however, makes it appear as though multiple are
// by clever timeslicing.

// The first goroutine immediately returns a 1 value on its channel while
// the second has an indefinite loop. With a single process, the program
// hangs, waiting for both goroutines to finish sequentially. With two
// processes, the program can exit as soon as the first returns.

// Run it like this:
// env GOMAXPROCS=1 go run concurrency-vs-parallelism.go
// env GOMAXPROCS=2 go run concurrency-vs-parallelism.go

package main

import "fmt"

func main() {
	c := make(chan int)

	finiteFunction := func() {
		c <- 1
	}

	indefiniteFunction := func() {
		for {
		}
		// Unreachable code.
		c <- 1
	}

	go finiteFunction()
	go indefiniteFunction()

	fmt.Println(<-c)
}
