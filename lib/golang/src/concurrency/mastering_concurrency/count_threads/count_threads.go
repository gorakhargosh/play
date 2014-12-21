package main

import (
	"fmt"
	"runtime"
)

func listThreads() int {
	threads := runtime.GOMAXPROCS(0)
	return threads
}

func main() {
	runtime.GOMAXPROCS(20)
	fmt.Printf("%d thread(s) available to Go.\n", listThreads())
}
