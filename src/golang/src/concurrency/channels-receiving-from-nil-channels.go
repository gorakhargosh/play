package main

import "fmt"

func main() {
	ch := make(chan bool, 2)
	ch <- true
	ch <- true

	close(ch)

	// Attempt to read from the channel more times
	// than the capacity of the channel.
	for i := 0; i < cap(ch)+1; i++ {
		v, ok := <-ch
		// Any reads after the channel is closed will receive
		// the zero value of the type of the channel and ok
		// will be set to false.
		fmt.Println(v, ok)
	}
}
