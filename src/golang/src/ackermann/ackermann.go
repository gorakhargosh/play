package main

import (
	"fmt"
	"time"
)

// ackermann computes ackermann computes ackermann.
func ackermann(m, n int) int {
	if m == 0 {
		return n + 1
	} else if n == 0 {
		return ackermann(m-1, 1)
	} else {
		return ackermann(m-1, ackermann(m, n-1))
	}
}

// Timeit times a given function and passes count as an argument to it.
func Timeit(label string, fn func() int) {
	start := time.Now()
	v := fn()
	elapsed := time.Since(start)
	fmt.Printf("%s = %d (%v)\n", label, v, elapsed)
}

func main() {
	for m := 0; m < 10; m++ {
		for n := 0; n < 10; n++ {
			Timeit(fmt.Sprintf("ackermann(%d,%d)", m, n), func() int {
				return ackermann(m, n)
			})
		}
	}
}
