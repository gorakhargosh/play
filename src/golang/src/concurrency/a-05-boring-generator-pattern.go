package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generates endless messages on a channel that it returns.
func boring(msg string) chan string {
	c := make(chan string)
	// The go routine is now embedded into the generator.
	// The computation or I/O is now neatly encapsulated by the
	// boring function and it simply returns a channel from which
	// a reader goroutine can read values.
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}
