package main

import "fmt"

func main() {
	f := 100.5
	var a int
	// a := int(100.5) // Generates an error.
	a = 100.5
	b := int(f) // Does not generate an error.
	fmt.Println(a, b)
}

// error text:
// type-conversion.go:7: constant 100.5 truncated to integer
