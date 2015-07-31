package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// rand.Seed(10)
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(1000))
	}
}
