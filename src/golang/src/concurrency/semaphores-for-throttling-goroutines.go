package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var running int64

func work() {
	atomic.AddInt64(&running, 1)
	fmt.Printf("[%d ", running)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	atomic.AddInt64(&running, -1)
	fmt.Print("]")
}

func worker(sema chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	// This semaphore blocks reading only while there are no "slots" available
	// for this worker depending upon its capacity.
	<-sema
	work()

	// Signal to the semaphore that we're done with work and a "slot" is
	// available for someone else to pick it up.
	sema <- true
}

func main() {
	var wg sync.WaitGroup

	// Semaphore implemented as a buffered channel (does non-blocking I/O
	// until capacity filled).
	sema := make(chan bool, 20)
	numWorkers := 1000
	// The wait group must be updated with its count before the goroutine is
	// scheduled to prevent race conditions.
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(sema, &wg)
	}

	// Make enough "slots" available on the semaphore.
	for i := 0; i < cap(sema); i++ {
		sema <- true
	}

	// Using a wait group is better than this guesstimate of a sleep.
	// time.Sleep(time.Duration(30) * time.Second)
	wg.Wait()

	fmt.Println()
}
