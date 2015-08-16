// Demonstrates communication between two goroutines over a channel. Instead of
// simply dumping values straight to standard output, the emitter sends values
// onto the channel, which the main goroutine reads from in order to display
// values on standard output.
//
// Communication over a regular channel is a synchronization operation because
// the send and the receive operations are in lockstep. Buffered channels don't
// synchronize, however. They're very much like mailboxes in Erlang.
//
// The Go approach to concurrent software:
//
// Don't communicate by sharing memory; share memory by communicating.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generates endless messages.
func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		// Send on the channel.
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go boring("boring!", c)
	// We're only going to read 5 values from the channel.
	for i := 0; i < 5; i++ {
		// Receiving from the channel.
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}
