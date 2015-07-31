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
	// This times out the entire conversation (applies to the loop).
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
}
