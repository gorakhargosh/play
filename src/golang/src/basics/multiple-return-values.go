package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// If we only want the second value.
	_, c := vals()
	fmt.Println(c)
}
