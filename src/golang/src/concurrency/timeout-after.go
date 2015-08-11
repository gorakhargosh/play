package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func gen(n int) <-chan int {
	c := make(chan int)
	go func() {
		time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
		c <- n
	}()
	return c
}

// show A OMIT

// Use time.After() instead. Only for demonstration. Actual implementation
// is different.
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
	ch := gen(150) // max ms
	select {       // HL
	case v := <-ch: // HL
		fmt.Println("success: read from channel:", v)
	case <-After(80 * time.Millisecond): // HL
		fmt.Println("error: channel timed out")
	} // HL
}

// end show A OMIT
