package main

import (
	"fmt"
	"time"
)

// boring just prints a message every second.
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}

func main() {
	boring("boring!")
}
