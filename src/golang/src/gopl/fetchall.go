// A program to concurrently fetch responses from multiple Websites.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for i, _ := range os.Args[1:] {
		fmt.Println(i)
		fmt.Println(<-ch)
	}
	fmt.Printf("total time elapsed: %v", time.Since(start).Seconds())
}

// fetch fetches a response given a URL and dumps it into a channel.
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	_, err = io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	ch <- fmt.Sprintf("elapsed: %v", time.Since(start).Seconds())
}
