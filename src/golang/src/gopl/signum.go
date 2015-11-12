package main

import "fmt"

func signum(num int) int {
	switch false {
	case num > 0:
		return 1
	default:
		return 0
	case num < 0:
		return -1
	}
}

func main() {
	fmt.Println(signum(-20), signum(0), signum(20))
}
