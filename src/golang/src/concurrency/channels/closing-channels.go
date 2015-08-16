package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Meow!"
		close(ch) // indicate end of communication // HL
		// ch <- "uncomment to panic!" // HL
	}()

	fmt.Println(<-ch) // blocking

	fmt.Println(<-ch) // non-blocking; receives zero-value // HL
	fmt.Println(<-ch) // non-blocking; receives zero-value // HL
	v, ok := <-ch     // v is "", ok is false // HL
	fmt.Printf("value %q from channel? %v", v, ok)
}
