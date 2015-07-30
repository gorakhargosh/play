// The unbuffered channel is not strictly required. All this achieves here is
// allow the send operation into the channel to be asynchronous. Also, as I am
// learning, using channels like gotos here can cause confusion because this
// uses a "flow-of-control" approach to synchronization rather than a
// "flow-of-data" approach.
//
// If you'd want to wait for goroutines to finish, using a WaitGroup is a much
// more readable option, because it makes the "waiting" explicit.
package main

import (
	"math/rand"
	"time"
)

// show A OMIT
func worker(done chan bool) {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	done <- true // HL
}

func main() {
	// This does not need to be a buffered channel.
	done := make(chan bool, 1) // HL
	go worker(done)
	<-done // HL
}

// end show A OMIT
