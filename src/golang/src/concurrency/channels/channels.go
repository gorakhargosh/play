package main

import "fmt"

func main() {
	messages := make(chan string) // HL
	go func() {
		messages <- "ping" // HL
	}()

	msg := <-messages // HL
	fmt.Println(msg)
}
