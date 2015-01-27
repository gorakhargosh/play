package main

import (
	"fmt"
	"runtime"
)

func showNumber(num int) {
	fmt.Println(num)
}

func main() {
	runtime.GOMAXPROCS(2)
	iterations := 10

	for i := 0; i <= iterations; i++ {
		go showNumber(i)
	}
	runtime.Gosched()
	fmt.Println("Goodbye!")
}
