package main

import "fmt"

func main() {
	aValue := new(int)

	readValue := func() {
		fmt.Println(*aValue)
	}

	defer readValue()

	for i := 0; i < 100; i++ {
		*aValue++
	}
}
