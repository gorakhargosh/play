// Round trip quit channels.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Multiplexes reading from multiple channels fanning into a single channel
// for the reader.
func fanIn(a, b <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case c <- <-a:
			case c <- <-b:
			}
		}
	}()
	return c
}

// cleanup pretends to perform some clean up.
func cleanup() {
	fmt.Println("Cleaned up!")
}

// Generates endless messages on a channel that it returns.
func boring(msg string, quit chan bool) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			case <-quit:
				// The other goroutine needs to wait for this case block to complete
				// executing too!
				cleanup()
				// Signal to the other waiting goroutine that we're done with the
				// clean up.
				quit <- true
				return
			}
		}
	}()
	return c
}

func main() {

	// Now we'll tell joe to shut up using the quit channel.
	quit := make(chan bool)
	joe := boring("Joe!", quit)

	// Seed the random number generator first.
	rand.Seed(time.Now().UTC().UnixNano())

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-joe)
	}
	quit <- true
	<-quit
}
