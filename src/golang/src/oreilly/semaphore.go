package main

import (
	"fmt"
	"math/rand"
	"time"
)

func work() {
	fmt.Print("[")
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Print("]")
}

func worker(sema chan bool) {
	<-sema
	work()
	sema <- true
}

func main() {
	sema := make(chan bool, 100)
	for i := 0; i < 1000; i++ {
		go worker(sema)
	}

	for i := 0; i < cap(sema); i++ {
		sema <- true
	}

	time.Sleep(time.Duration(30) * time.Second)
	fmt.Println()
}
