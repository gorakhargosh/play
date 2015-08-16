package main

import (
	"fmt"
	"math/rand"
	"time"
)

// show A OMIT
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	// Like running a bash shell command in the background using `&'.
	// Example: python -m SimpleHTTPServer &
	go boring("boring!") // HL

	// main returns immediately as we don't wait for the goroutine above to
	// complete.
}

// end show A OMIT
