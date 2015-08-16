package main

import (
	"fmt"
	"time"
)

// Ball represents a ball that keeps count of the number of times it has been
// hit.

// show A OMIT
type Ball struct{ hits int } // HL

func player(name string, table chan *Ball) {
	for {
		ball := <-table // HL
		ball.hits++     // HL
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball // HL
	}
}

func main() {
	table := make(chan *Ball) // HL
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball) // comment for deadlock.  // HL
	time.Sleep(1 * time.Second)
	<-table // game over; grab the ball. // HL
}

// end show A OMIT
