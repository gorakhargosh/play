package main

import "fmt"

func main() {
	// defer stacks up calls and calls are executed in LIFO.
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
