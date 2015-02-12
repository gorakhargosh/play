package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	interrupt := func(seconds int, timeout_chan chan int) {
		time.Sleep(time.Second * time.Duration(seconds))
		timeout_chan <- seconds
	}

	go interrupt(1, c1)
	go interrupt(2, c2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("recieved", msg2)
		}
	}
}
