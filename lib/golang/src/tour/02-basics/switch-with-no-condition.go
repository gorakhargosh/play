package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// Long if-then-else chains can be written using this.
	switch true {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
