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
	boring("boring!")
}

// end show A OMIT
