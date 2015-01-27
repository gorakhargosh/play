package main

import "os"

func main() {
	file, _ := os.Create("/defer.txt")
	defer file.Close()

	// No op.
	for {
		break
	}
}
