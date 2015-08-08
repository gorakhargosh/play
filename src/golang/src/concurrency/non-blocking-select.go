package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		time.Sleep(1 * time.Second) // HL
		c <- 1                      // HL
	}()

	select { // HL
	case v := <-c:
		fmt.Printf("We received the value %q.", v) // Never displayed.
	default: // HL
		fmt.Println("non-blocking select moves on; does not wait.") // HL
	} // HL
	fmt.Println("Done.")
}
