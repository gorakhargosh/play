package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var whatHrSaid bool

const (
	Day   time.Duration = time.Hour * 24
	Month               = Day * 30
)

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.BoolVar(&whatHrSaid, "kya", true, "what hr says")
}

func work() {
	time.Sleep(time.Duration(rand.Intn(8)) * time.Millisecond) // Read hours.
	fmt.Println("work work work work")
}

func helloWorld() {
	fmt.Println("hello world!")
}

func main() {
	hr := make(chan bool)
	noticePeriod := 2 * Month

	timeout := time.After(noticePeriod)

	flag.Parse()

	go func() {
		// HR.
		time.Sleep(time.Duration(10) * time.Millisecond) // Read hours.
		hr <- whatHrSaid
	}()

	for {
		select {
		case response := <-hr:
			if response {
				work()
				return
			} else {
				helloWorld()
			}
			return
		case <-timeout:
			helloWorld()
			return
		default:
			work()
		}
	}
}
