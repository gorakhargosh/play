package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

func searchEngine(kind string) Search {
	return func(query string) Result {
		// Pause for a while and return a string result.
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

var (
	Web1   = searchEngine("web1")
	Web2   = searchEngine("web2")
	Web3   = searchEngine("web3")
	Image1 = searchEngine("image1")
	Image2 = searchEngine("image2")
	Image3 = searchEngine("image3")
	Video1 = searchEngine("video1")
	Video2 = searchEngine("video2")
	Video3 = searchEngine("video3")
)

type GoogleFunc func(query string) (results []Result)

// Serial queries.
func Google1(query string) (results []Result) {
	results = append(results, Web1(query))
	results = append(results, Image1(query))
	results = append(results, Video1(query))
	return
}

// Parallel queries.
func Google2(query string) (results []Result) {
	c := make(chan Result)
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

// Parallel and time-bound.
func Google2P1(query string) (results []Result) {
	c := make(chan Result)
	numbackends := 3
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
	for i := 0; i < numbackends; i++ {
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

// Final replicated, parallel, and time-bound.
func Google3(query string) (results []Result) {
	c := make(chan Result)
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

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func TimeQuery(query string, search GoogleFunc, label string) {
	start := time.Now()
	results := search(query)
	elapsed := time.Since(start)
	fmt.Println(label, results)
	fmt.Println(label, elapsed)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	TimeQuery("golang", Google1, "Google1")
	TimeQuery("golang", Google2, "Google2")
	TimeQuery("golang", Google2P1, "Google2.1")
	TimeQuery("golang", Google3, "Google3")
}
