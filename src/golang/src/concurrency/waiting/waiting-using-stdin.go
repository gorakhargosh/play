package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// show A OMIT
func main() {
	go boring("message")

	var discard string
	fmt.Scanf("%s", &discard) // HL
	fmt.Printf("You entered: %q; exiting.\n", discard)
}

// end show A OMIT
