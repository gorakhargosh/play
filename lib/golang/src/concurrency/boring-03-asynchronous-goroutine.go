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

	// Aha, but then you just exited the process and boring is no longer doing
	// anything because the main goroutine (analogous to thread) exited. We need
	// to wait here somehow.
}
