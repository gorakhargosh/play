// This example demonstrates the proper use of a wait-group
// to wait for all concurrent I/O to complete without causing
// race conditions in the wait-group counter.
//
// See: http://stackoverflow.com/questions/19208725/example-for-sync-waitgroup-correct
//
// Waitgroups panic if the counter falls below zero. The counter starts
// at zero, each Done() is a -1 and each Add() depends on the parameter.
// So, you need the Add() to be guaranteed to come before the Done() to
// avoid panics.

// In Go, such guarantees are given by the memory model.

// The memory model states that all statements in a single goroutine
// appear to be executed in the same order as they are written. It is
// possible that they won't actually be in that order, but the outcome
// will be as if it was. It is also guaranteed that a goroutine
// doesn't run until after the go statement that calls it. Since
// the Add() occurs before the go statement and the go statement
// occurs before the Done(), we know the Add() occurs before the Done().

// If you were have the go statement come before the Add(), the
// program may operate correctly. However, it would be a race condition
// because it would not be guaranteed.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func doWork(i int, wg *sync.WaitGroup) {
	// wg.Done() essentially subtracts 1 from the counter.
	// Therefore, you should ensure that you're incrementing the
	// counter by only 1 unit per goroutine, not more.
	defer wg.Done()
	// Pretend to do some blocking I/O.
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	fmt.Printf("Done with work: %d\n", i)
}

func main() {
	// The zero-value for this is sufficient.
	var wg sync.WaitGroup
	count := 4

	// Calling wg.Add(1) repeatedly is non-pragmatic if you know the count.
	wg.Add(count)
	for i := 0; i < count; i++ {
		// The wait group count must be updated before scheduling the goroutine
		// for execution to prevent race conditions maintaining the counter (
		// wg.Done() would drop the counter by 1 and the wait group can panic
		// when the counter goes below 0.
		//

		// wg.Add(1)
		go doWork(i, &wg)
	}

	// This will block until the counter drops to 0.
	wg.Wait()

	fmt.Println("We waited until the last piece of work was complete. Done.")
}
