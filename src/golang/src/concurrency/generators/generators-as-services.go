package main

import (
	"fmt"
	"math/rand"
	"time"
)

// show A OMIT
func boring(msg string) chan string {
	c := make(chan string) // HL
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) // HL
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // HL
}

func main() {
	joe := boring("Joe!") // HL
	ann := boring("Ann!") // HL
	for i := 0; i < 10; i++ {
		fmt.Println(<-joe) // HL
		fmt.Println(<-ann) // Reading from ann blocks until joe has been read from. // HL
	}
	fmt.Println("You're boring; I'm leaving.")
}

// end show A OMIT
