package main

import "fmt"

func main() {
	messages := make(chan string, 2) // HL

	messages <- "buffered" // non-blocking // HL
	messages <- "channel"  // non-blocking // HL
	// messages <- "deadlock" // HL

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
