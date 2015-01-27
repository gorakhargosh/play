package main

import (
	"fmt"
	"time"
)

func outputText(text string, count int) {
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(text)
	}
}

func main() {
	go outputText("hello", 3)
	outputText("world", 5)
}
