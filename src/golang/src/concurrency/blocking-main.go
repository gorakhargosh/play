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

		// We can use a fixed interval.
		// time.Sleep(time.Second) // HL

		// However, we will use a random time interval to better simulate real
		// blocking computation or I/O.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond) // HL
	}
}

func main() {
	// Blocks the main goroutine.
	boring("boring!") // HL
	// We won't reach here until boring is done.
}

// end show A OMIT
