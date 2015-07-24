package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var running int64 = 0

func work() {
	atomic.AddInt64(&running, 1)
	fmt.Printf("[%d ", running)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	atomic.AddInt64(&running, -1)
	fmt.Print("]")
}

func worker(sema chan bool) {
	<-sema
	work()
	sema <- true
}

func main() {
	sema := make(chan bool, 20)
	for i := 0; i < 1000; i++ {
		go worker(sema)
	}

	for i := 0; i < cap(sema); i++ {
		sema <- true
	}

	time.Sleep(time.Duration(30) * time.Second)
	fmt.Println()
}
