package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		// We introduce a bit of randomness to make the messages print a little less
		// deterministically.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	// Do the boring thing in the background.
	go boring("boring!")

	// We wait for the boring goroutine to do its stuff... for 2 seconds only.
	// This is a very bad way of estimating how long to wait for a task to
	// complete because we can overestimate and underestimate, and rarely ever, be
	// precise.

	// How can you be sure that the boring function will take 2 seconds to
	// complete? I could have a fast computer, or a slow computer, or the function
	// could be doing I/O and waiting for it to complete, or it may just do
	// nothing and exit; in which case, attempting to estimate the amount of time
	// required for a particular task to complete is futile and pointless. We'll
	// look at better ways to handle this situation soon.
	fmt.Println("I'm listening")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}
