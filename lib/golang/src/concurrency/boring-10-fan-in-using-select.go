// Don't communicate by sharing memory; share memory by communicating.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// fanIn is another generator that does:
// https://talks.golang.org/2012/concurrency/images/gophermegaphones.jpg
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	// We don't need to create multiple goroutines to fan in multiple channels. We
	// can use a select for exactly that--selecting one from many possible
	// communications.
	go func() {
		for {
			select {
			case c <- <-input1:
			case c <- <-input2:
			}
		}
	}()
	return c
}

func boring(msg string) <-chan string {
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
	c := fanIn(boring("Joe!"), boring("Ann!"))

	// We now need to wait for 10 iterations instead of 5.
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You're boring; I'm leaving.")
}
