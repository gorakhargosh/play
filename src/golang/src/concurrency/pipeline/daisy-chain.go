package main

import "fmt"

// |o.O|P <- |o.O|P <- |o.O|P <- |O.O|

// Chinese whispers (mutation of the content).
// Simply counts the number of I/O whispers.
func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 100000
	leftmost := make(chan int)

	right := leftmost // Just to initialize.
	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		// Schedules communication on all of these channels
		// and they are all blocked on receive until the send operation.
		go f(left, right)
		left = right
	}

	go func(c chan int) {
		// Now send on the rightmost channel. Still blocking.
		c <- 1
	}(right)

	// Unblock the communication because the reader is now ready.
	fmt.Println(<-leftmost)
}
