package main

import "fmt"

func main() {
	c := make(chan int, 2)
	go func() {
		c <- 1
		c <- 2
	}()

	fmt.Println(<-c) // 1
	close(c)         // Indicate the closing of the channel. // HL
	fmt.Println(<-c) // 2 // HL
	fmt.Println(<-c) // zero value. // HL
}
