package main

import "fmt"

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

func main() {
	for m := 0; m < 10; m++ {
		for n := 0; n < 10; n++ {
			fmt.Printf("ackermann(%d,%d) = %d\n", m, n, ackermann(m, n))
		}
	}
}
