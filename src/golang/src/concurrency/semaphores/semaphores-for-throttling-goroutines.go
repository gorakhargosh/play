package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// show atomicCounting OMIT
var running int64 // HL

func count() {
	atomic.AddInt64(&running, 1) // HL
	fmt.Printf("[%d ", running)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	atomic.AddInt64(&running, -1) // HL
	fmt.Print("]")
}

// end show atomicCounting OMIT

func worker(sema chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	<-sema // Wait until "slot" available. // HL
	count()
	sema <- true // Signal "slot" available. // HL
}

func main() {
	var wg sync.WaitGroup

	sema := make(chan bool, 100) // non-blocking write until capacity // HL
	numWorkers := 1000           // HL
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(sema, &wg) // HL
	}

	// Make enough "slots" available on the semaphore.
	for i := 0; i < cap(sema); i++ { // HL
		sema <- true // HL
	} // HL
	wg.Wait()
	fmt.Println("Done!")
}
