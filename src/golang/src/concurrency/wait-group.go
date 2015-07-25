// This example demonstrates the proper use of a wait-group
// to wait for all concurrent I/O to complete without causing
// race conditions in the wait-group counter.
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

	for i := 0; i < 4; i++ {
		// The wait group count must be updated before scheduling the goroutine
		// for execution to prevent race conditions maintaining the counter (
		// wg.Done() would drop the counter by 1 and the wait group can panic
		// when the counter goes below 0.
		wg.Add(1)
		go doWork(i, &wg)
	}

	// This will block until the counter drops to 0.
	wg.Wait()

	fmt.Println("We waited until the last piece of work was complete. Done.")
}
