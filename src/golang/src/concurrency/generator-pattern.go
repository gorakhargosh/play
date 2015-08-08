package main

import (
	"fmt"
	"math/rand"
	"time"
)

// show A OMIT
// Generates endless messages on a channel that it returns.
func boring(msg string) chan string {
	c := make(chan string) // HL
	// The go routine is now embedded into the generator.
	// The computation or I/O is now neatly encapsulated by the
	// boring function and it simply returns a channel from which
	// a reader goroutine can read values.
	go func() { // HL
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) // HL
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}() // HL
	return c // HL
}

func main() {
	c := boring("boring!") // HL
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // HL
	}
	fmt.Println("You're boring; I'm leaving.")
}

// end show A OMIT
