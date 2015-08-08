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

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func consume(who chan string) {
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-who)
	}
}

// show A OMIT
func boring(msg string, quit chan bool) chan string { // HL
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			case <-quit: // HL
				return // HL
			}
		}
	}()
	return c
}

func main() {
	quit := make(chan bool) // HL
	consume(boring("Joe!", quit))
	quit <- true // HL
}

// end show A OMIT
