package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func boring(msg string, count int, wg *sync.WaitGroup) {
	for i := 0; i < count; i++ {
		fmt.Println(msg, i)
		// We introduce a bit of randomness to make the messages print a little less
		// deterministically.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
	wg.Done()
}

func main() {
	// Do the boring thing in the background.
	wg := new(sync.WaitGroup)

	go boring("boring!", 10, wg)

	// Now the main goroutine waits precisely for as much time as is required for
	// the boring goroutine to complete.
	fmt.Println("I'm listening")
	wg.Add(1)
	wg.Wait()
}
