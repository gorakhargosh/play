package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Multiplexes reading from multiple channels fanning into a single channel
// for the reader.
func fanIn(a, b <-chan string) <-chan string {
	c := make(chan string)

	// We replace these two goroutines with a single one that uses select
	// to choose one of the communications that is ready every iteration
	// of the endless loop.
	//
	// go func() {
	// 	for {
	// 		c <- <-a
	// 	}
	// }()
	// go func() {
	// 	for {
	// 		c <- <-b
	// 	}
	// }()

	go func() {
		for {
			select {
			case c <- <-a:
			case c <- <-b:
			}
		}
	}()
	return c
}

// Generates endless messages on a channel that it returns.
func boring(msg string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	joe := boring("Joe!")
	ann := boring("Ann!")
	c := fanIn(joe, ann)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}
