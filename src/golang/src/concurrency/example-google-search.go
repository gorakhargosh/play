package main

import (
	"fmt"
	"math/rand"
	"time"
)

// show fakeEngines OMIT
var (
	Web1   = fakeSearch("web1") // HL
	Web2   = fakeSearch("web2")
	Web3   = fakeSearch("web3")
	Image1 = fakeSearch("image1") // HL
	Image2 = fakeSearch("image2")
	Image3 = fakeSearch("image3")
	Video1 = fakeSearch("video1") // HL
	Video2 = fakeSearch("video2")
	Video3 = fakeSearch("video3")
)

// end show fakeEngines OMIT

type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

type SearchFunc func(query string) []Result

// Sequential.
func Google1(query string) (results []Result) {
	results = append(results, Web1(query))   // HL
	results = append(results, Image1(query)) // HL
	results = append(results, Video1(query)) // HL
	return results
}

// Concurrent.
func Google2(query string) (results []Result) {
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
func Google3(query string) (results []Result) {
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

// Uses the first response from the replicas.
func First(query string, replicas ...Search) Result { // varargs // HL
	c := make(chan Result)
	for index := range replicas { // HL
		go func(i int) { // HL
			c <- replicas[i](query) // HL
		}(index) // HL
	} // HL
	return <-c
}

// Concurrent, time-boxed, and replicated.
func Google4(query string) (results []Result) {
	c := make(chan Result)

	// Fan-in pattern.
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
	Timeit("1: sequential", Google1, query)
	Timeit("2: concurrent", Google2, query)
	Timeit("3: concurrent, time-bound", Google3, query)
	Timeit("4: concurrent, time-bound, replicated", Google4, query)
}
