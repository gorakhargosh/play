package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/net/context"
)

// show fakeEngines OMIT
var (
	Web1   = makeBackend("web1") // HL
	Web2   = makeBackend("web2")
	Web3   = makeBackend("web3")
	Image1 = makeBackend("image1") // HL
	Image2 = makeBackend("image2")
	Image3 = makeBackend("image3")
	Video1 = makeBackend("video1") // HL
	Video2 = makeBackend("video2")
	Video3 = makeBackend("video3")
)

// end show fakeEngines OMIT

var (
	CWeb1   = makeCancelableBackend("c_web1") // HL
	CWeb2   = makeCancelableBackend("c_web2")
	CWeb3   = makeCancelableBackend("c_web3")
	CImage1 = makeCancelableBackend("c_image1") // HL
	CImage2 = makeCancelableBackend("c_image2")
	CImage3 = makeCancelableBackend("c_image3")
	CVideo1 = makeCancelableBackend("c_video1") // HL
	CVideo2 = makeCancelableBackend("c_video2")
	CVideo3 = makeCancelableBackend("c_video3")
)

// show firstComeFirstServed OMIT
type Result struct {
	Err   error
	Value string
}                                     // HL
type Search func(query string) Result // HL

// Uses the first response from the replicas.
func First(query string, replicas ...Search) Result { // HL
	// Buffered channel to hold obtained results.
	c := make(chan Result, len(replicas))                  // HL
	search := func(replica Search) { c <- replica(query) } // HL
	for _, replica := range replicas {
		go search(replica) // HL
	}
	return <-c // The value is returned, not the channel. // HL
}

// end show firstComeFirstServed OMIT

type CancelableSearch func(ctx context.Context, query string) Result // HL

func CancelableFirst(ctx context.Context, query string, replicas ...CancelableSearch) Result { // HL
	c := make(chan Result, len(replicas))
	ctx, cancel := context.WithCancel(ctx)                                // HL
	defer cancel()                                                        // HL
	search := func(replica CancelableSearch) { c <- replica(ctx, query) } // HL
	for _, replica := range replicas {
		go search(replica)
	}
	select { // HL
	case <-ctx.Done(): // HL
		return Result{Err: ctx.Err()} // HL
	case r := <-c: // HL
		return r // HL
	}
}

func makeBackend(kind string) Search { // HL
	return func(query string) Result { // HL
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{Value: fmt.Sprintf("%s result for %q\n", kind, query)}
	}
}

func makeCancelableBackend(kind string) CancelableSearch { // HL
	return func(ctx context.Context, query string) Result { // HL
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{
			Err:   ctx.Err(),
			Value: fmt.Sprintf("%s result for %q\n", kind, query),
		}
	}
}

type SearchFunc func(query string) []Result

// Sequential.
func Google1(query string) (results []Result) { // HL
	results = append(results, Web1(query))   // HL
	results = append(results, Image1(query)) // HL
	results = append(results, Video1(query)) // HL
	return results
}

// Concurrent.
func Google2(query string) (results []Result) { // HL
	c := make(chan Result)

	// Fan-in (multiplexing) pattern.
	go func() { c <- Web1(query) }()   // HL
	go func() { c <- Image1(query) }() // HL
	go func() { c <- Video1(query) }() // HL

	for i := 0; i < 3; i++ {
		result := <-c // HL
		results = append(results, result)
	}
	return
}

// Concurrent and time-bound.
func Google3(query string) (results []Result) { // HL
	c := make(chan Result)

	go func() { c <- Web1(query) }()
	go func() { c <- Image1(query) }()
	go func() { c <- Video1(query) }()

	timeout := time.After(80 * time.Millisecond) // HL
	for i := 0; i < 3; i++ {
		select { // HL
		case result := <-c: // HL
			results = append(results, result)
		case <-timeout: // HL
			fmt.Println("timed out")
			return
		} // HL
	}
	return
}

// Concurrent, time-boxed, and replicated.
func Google4(query string) (results []Result) { // HL
	c := make(chan Result)

	// Fan-in and replication.
	go func() { c <- First(query, Web1, Web2, Web3) }()       // HL
	go func() { c <- First(query, Image1, Image2, Image3) }() // HL
	go func() { c <- First(query, Video1, Video2, Video3) }() // HL

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

// Concurrent, time-bound, replicated, and tail-latency reduction.
func Google5(query string) (results []Result) { // HL
	c := make(chan Result)
	ctx, cancel := context.WithCancel(context.Background())                     // HL
	defer cancel()                                                              // HL
	go func() { c <- CancelableFirst(ctx, query, CWeb1, CWeb2, CWeb3) }()       // HL
	go func() { c <- CancelableFirst(ctx, query, CImage1, CImage2, CImage3) }() // HL
	go func() { c <- CancelableFirst(ctx, query, CVideo1, CVideo2, CVideo3) }() // HL
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func Timeit(label string, fn SearchFunc, query string) []Result {
	start := time.Now()
	results := fn(query)
	elapsed := time.Since(start)
	fmt.Printf("%s: %s\n%v\n\n", label, elapsed, results)
	return results
}

func main() {
	rand.Seed(time.Now().UnixNano())
	query := "golang"
	Timeit("Google1", Google1, query)
	Timeit("Google2", Google2, query)
	Timeit("Google3", Google3, query)
	Timeit("Google4", Google4, query)
	Timeit("Google5", Google5, query)
}
