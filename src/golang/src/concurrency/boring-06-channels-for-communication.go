package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		// Expression to be sent can be any suitable value.
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	// A channel connects the main and boring goroutines so that they can
	// communicate.
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		// Receive expression is just a value.
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}
