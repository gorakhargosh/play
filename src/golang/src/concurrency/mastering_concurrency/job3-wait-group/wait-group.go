package main

import (
	"fmt"
	"sync"
	"time"
)

func outputText(text string, count int, g *sync.WaitGroup) {
	defer g.Done()
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(text)
	}
}

func main() {
	g := new(sync.WaitGroup)
	go outputText("hello", 3, g)
	go outputText("world", 5, g)
	g.Add(2)
	g.Wait()
}
