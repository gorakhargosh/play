package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func gen(n int) <-chan time.Duration { // HL
	c := make(chan time.Duration) // HL
	go func() {
		d := time.Duration(rand.Intn(n)) * time.Millisecond
		time.Sleep(d)
		c <- d // HL
	}()
	return c
}

// show A OMIT

// Use time.After() instead. Only for demonstration. Actual implementation differs.
func After(t time.Duration) <-chan bool { // HL
	// buffer = 1; goroutine won't hang around forever if alternative
	// communication happens before the timeout is reached. The timeout channel
	// will eventually be deallocated by the garbage collector.
	timeout := make(chan bool, 1) // HL
	go func() {
		time.Sleep(t)
		timeout <- true // non-blocking send will definitely succeed // HL
	}()
	return timeout
}

func main() {
	ch := gen(150)                // max ms
	tmax := 80 * time.Millisecond // ms
	select {                      // HL
	case v := <-ch: // HL
		fmt.Printf("success: read from channel: %v\n", v)
	case <-After(tmax): // HL
		fmt.Printf("error: channel timed out (%v)\n", tmax)
	} // HL
}

// end show A OMIT
