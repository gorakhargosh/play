package main

import (
	"fmt"
	"os"
)

const (
	message = "%d %d\n"
	answer1 = iota
	answer2
)

func Foo(message string) {
	fmt.Println(message)
}

func main() {
	pi := float64(3.14)
	nine := uint(9)
	isTrue := !false
	var is bool

	fmt.Printf("Value: %.2f\n", pi)
	fmt.Printf("Value: %v, Type: %T\n", nine, nine)
	fmt.Printf("Value: %v, Truth value: %t\n", isTrue, isTrue)
	fmt.Println(is)
	fmt.Printf("Value: %x %x\n", byte(10), byte(100))

	atoz := `the quick brown fox jumps over the lazy dog\n`
	fmt.Printf("%s\n", atoz[16:])
	fmt.Printf("%d\n", len(atoz))

	if n, err := fmt.Printf("Hello, world!\n"); err != nil {
		os.Exit(1)
	} else {
		fmt.Printf("Printed %d bytes\n", n)
	}
}
