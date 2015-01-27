package main

import (
	"fmt"
	"math/rand"
	"time"
)

// boring just prints a message after random intervals.
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		// We introduce a bit of randomness to make the messages print a little less
		// deterministically.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	boring("boring!")
}
