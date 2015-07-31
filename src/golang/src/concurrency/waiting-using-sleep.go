package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// show A OMIT
func main() {
	go boring("boring!")

	// We wait 2 seconds. There's no communication between this main goroutine and
	// the boring goroutine.
	fmt.Println("I'm listening")
	time.Sleep(2 * time.Second) // HL
	fmt.Println("You're boring; I'm leaving.")
}

// end show A OMIT
