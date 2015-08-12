package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Meow!"
		close(ch) // indicate end of communication // HL
	}()
	fmt.Println(<-ch)

	// All following receive operations are non-blocking and receive the nil value
	// for the channel type.
	fmt.Println(<-ch) // HL
	fmt.Println(<-ch) // HL
	v, ok := <-ch     // v is "", ok is false // HL
	fmt.Printf("value %q from channel? %v", v, ok)
}
