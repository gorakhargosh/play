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

func main() {
	joe := boring("Joe!")
	ann := boring("Ann!")
	for i := 0; i < 10; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann) // Reading from ann blocks until joe has been read from.
	}
	fmt.Println("You're boring; I'm leaving.")
}
