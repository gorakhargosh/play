package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	fmt.Printf("Hello, world!\n")
	s := fmt.Sprintf("Hello, %s", "world")
	fmt.Printf("%s!\n", s)
}
