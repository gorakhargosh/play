package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

// show A OMIT
func fanIn(a, b <-chan string) <-chan string { // HL
	c := make(chan string)
	go func() { // HL
		for { // HL
			select { // HL
			case c <- <-a: // HL
			case c <- <-b: // HL
			} // HL
		} // HL
	}() // HL
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

// end show A OMIT
