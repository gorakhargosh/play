package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		// Fixed interval.
		// time.Sleep(time.Second)

		// Random sleep interval.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	// Ignore the boring thing that takes a lot of time.
	// Like running a shell command in the backgroun dusing the ampersand.
	go boring("boring!")

	// We wait for the boring function to complete for 2 seconds.
	fmt.Println("I'm listening")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}
