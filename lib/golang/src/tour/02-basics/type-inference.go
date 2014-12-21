package main

import "fmt"

func main() {
	// The value of the variable determines its type.
	var i int
	v := 42 // change me!
	v := i
	fmt.Printf("v is of type %T\n", v)
}
