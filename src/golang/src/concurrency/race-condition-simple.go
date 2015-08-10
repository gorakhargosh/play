package main

import "fmt"

func main() {
	wait := make(chan bool)

	n := 0 // HL
	go func() {
		n++         // One access: read, increment, write. // HL
		close(wait) // Resume any blocking reads from wait.
	}()
	n++ // Another access: read, increment, write. // HL

	<-wait
	fmt.Println(n) // Unspecified output. // HL
}
