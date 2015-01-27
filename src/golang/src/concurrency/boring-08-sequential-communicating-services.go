// Don't communicate by sharing memory; share memory by communicating.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator: function that returns a channel and is asynchronous by design.

// boring now returns a channel and acts like a "generator".
func boring(msg string) <-chan string {
	c := make(chan string)

	// Hello immediate functions, nested functions, and closures. We launch the
	// goroutine from inside the boring function to "hide" or "conceal" the
	// concurrent aspect of the boring function from the caller; this keeps their
	// lives simpler, while we do the heavy lifting.
	go func() {
		for i := 0; ; i++ {
			// Expression to be sent can be any suitable value.
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	// Return the channel to the caller.
	return c
}

func main() {
	// Using boring as a service.
	joe := boring("Joe!")
	ann := boring("Ann!")

	for i := 0; i < 5; i++ {
		// First we read from joe and keep ann waiting until joe has generated a
		// value. Ann is kept waiting even if she can generate faster... not a very
		// good approach. Ann is more talkative.
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're boring; I'm leaving.")
}
