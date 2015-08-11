// The following is a three-stage pipeline with cancellation. That is, a
// downstream stage can signal all upstream stages to stop sending data down the
// pipes.
//
// Here are the guidelines for a pipeline construction:
//
// 1. emitter stages close their outbound channels when all the send operations
//    are done.
//
// 2. emitter stages also select between outbound channels and done channels.
//
// 3. emitter stages exit early from the done case block and use defer to close
//    outbound channels.
//
// 4. stages keep receiving values from inbound channels until those channels
//    are closed or the senders are unblocked.
//
// We need a way to tell an unknown and unbounded number of goroutines to stop
// sending their values downstream. In Go, we can do this by closing a channel,
// because a receive operation on a closed channel can always proceed
// immediately, yielding the element type's zero value.

// This means that main can unblock all the senders simply by closing the done
// channel. This close is effectively a broadcast signal to the senders. We
// extend each of our pipeline functions to accept done as a parameter and
// arrange for the close to happen via a defer statement, so that all return
// paths from main will signal the pipeline stages to exit.
//
// More information at https://blog.golang.org/pipelines
package main

import (
	"fmt"
	"sync"
)

// gen emits the specified numbers on a channel that it returns. Closing the
// done signal channel causes gen to stop emitting more numbers.
func gen(done chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// sq reads numbers from a channel and emits the square of each on a channel
// that it returns. Closing the done channel causes sq to stop emitting numbers.
func sq(done chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// merge fans in the values from multiple channels and also allows downstream to
// indicate that no more values need to be emitted allowing for cancelation.
func merge(done chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs. output copies
	// values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are done.
	// This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	done := make(chan struct{})
	defer close(done)
	in := gen(done, 1, 2, 3, 4)
	a := sq(done, in)
	b := sq(done, in)

	out := merge(done, a, b)
	// We're only ever going to read one value, so we need to indicate upstream to
	// stop sending. We do that by closing the done channel. The main routine
	// effectively unblocks all the senders (the senders' selects alternate on the
	// done and the outbound channels).
	fmt.Println(<-out)
}
