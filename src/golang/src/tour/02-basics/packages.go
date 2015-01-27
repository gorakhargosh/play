package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// The random number generator is not seeded in this example.
	fmt.Println("My favorite number is", rand.Intn(10))
}
