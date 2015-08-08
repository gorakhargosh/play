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
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// show A OMIT
func worker(done chan bool) { // HL
	d := rand.Intn(1e3)
	time.Sleep(time.Duration(d) * time.Millisecond)
	fmt.Printf("Work took %d ms.\n", d)
	done <- true // HL
}

func main() {
	// This does not need to be a buffered channel.
	done := make(chan bool, 1) // HL
	go worker(done)
	<-done // HL
}

// end show A OMIT
