package main

import "fmt"

func main() {
	messages := make(chan string, 2) // HL
	messages <- "buffered"
	messages <- "channel"
	// messages <- "deadlock" // HL

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
