package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getPageLength(url string) (int, error) {
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}
	return len(body), nil
}

func getter(url string, info chan string) {
	length, err := getPageLength(url)
	if err == nil {
		info <- fmt.Sprintf("url %s; length: %d\n", url, length)
	}
}

func main() {
	var urls = []string{
		"http://www.google.com",
		"http://bbc.co.uk",
		"http://www.bing.com",
		"http://www.yahoo.com",
	}

	info := make(chan string)

	for _, url := range urls {
		go getter(url, info)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Print(<-info)
	}
}
