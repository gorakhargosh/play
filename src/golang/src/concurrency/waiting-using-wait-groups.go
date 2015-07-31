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

// show A OMIT
func doWork(i int, wg *sync.WaitGroup) {
	// wg.Done() essentially subtracts 1 from the counter.
	defer wg.Done() // HL
	d := rand.Intn(5)
	time.Sleep(time.Duration(d) * time.Second)
	fmt.Printf("Done with work %d in %d seconds\n", i, d)
}

func main() {
	var wg sync.WaitGroup // HL
	count := 4

	// Calling wg.Add(1) repeatedly is non-pragmatic if you know the count.
	wg.Add(count) // HL
	for i := 0; i < count; i++ {
		// The wait group count must be updated before scheduling the goroutine
		// for execution to prevent race conditions maintaining the counter.
		// wg.Add(1) // HL
		go doWork(i, &wg)
	}

	// This will block until the counter drops to 0.
	wg.Wait() // HL
}

// end show A OMIT
