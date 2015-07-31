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
	// Like running a shell command in the background using `&'.
	go boring("boring!") // HL

	// main immediately returns because we never waited for the goroutine
	// above to complete.
}
// end show A OMIT
