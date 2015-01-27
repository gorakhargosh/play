package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func outputText(text string, count int) {
	fileName := text + ".text"
	fileContents := ""
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Millisecond)
		fileContents += text
		fmt.Println(text)
	}
	err := ioutil.WriteFile(fileName, []byte(fileContents), 0644)
	if err != nil {
		panic("something went awry")
	}
}

func main() {
	go outputText("hello", 3)
	go outputText("world", 5)
}
