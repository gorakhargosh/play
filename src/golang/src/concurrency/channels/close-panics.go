package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	// var nc chan int // nil channel

	go func() {
		ch <- 1
		ch <- 2
	}()
	fmt.Println(<-ch)
	close(ch)

	// Uncomment any of the following to cause panics.
	// close(ch) // attempting to close a closed channel. // HL
	// close(nc) // attempting to close a nil channel. // HL
}
