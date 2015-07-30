package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web1   = fakeSearch("web1")
	Web2   = fakeSearch("web2")
	Web3   = fakeSearch("web3")
	Image1 = fakeSearch("image1")
	Image2 = fakeSearch("image2")
	Image3 = fakeSearch("image3")
	Video1 = fakeSearch("video1")
	Video2 = fakeSearch("video2")
	Video3 = fakeSearch("video3")
)

type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

type SearchFunc func(query string) []Result

// Serial.
func Google1(query string) (results []Result) {
	results = append(results, Web1(query))
	results = append(results, Image1(query))
	results = append(results, Video1(query))
	return results
}

// Concurrent.
func Google2(query string) (results []Result) {
	c := make(chan Result)

	// Fan-in pattern.
	go func() {
		c <- Web1(query)
	}()
	go func() {
		c <- Image1(query)
	}()
	go func() {
		c <- Video1(query)
	}()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

// Concurrent and time-boxed.
func Google3(query string) (results []Result) {
	c := make(chan Result)

	// Fan-in pattern.
	go func() {
		c <- Web1(query)
	}()
	go func() {
		c <- Image1(query)
	}()
	go func() {
		c <- Video1(query)
	}()

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

// Uses the first response from the replicas.
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	for index := range replicas {
		go func(i int) {
			c <- replicas[i](query)
		}(index)
	}
	return <-c
}

// Concurrent, time-boxed, and replicated.
func Google4(query string) (results []Result) {
	c := make(chan Result)

	// Fan-in pattern.
	go func() {
		c <- First(query, Web1, Web2, Web3)
	}()
	go func() {
		c <- First(query, Image1, Image2, Image3)
	}()
	go func() {
		c <- First(query, Video1, Video2, Video3)
	}()

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
	fmt.Printf("%s: %s\n%v\n", label, elapsed, results)
	return results
}

func main() {
	rand.Seed(time.Now().UnixNano())
	query := "golang"
	Timeit("serial", Google1, query)
	Timeit("concurrent", Google2, query)
	Timeit("concurrent, time-boxed", Google3, query)
	Timeit("concurrent, time-boxed, replicated", Google4, query)
}
