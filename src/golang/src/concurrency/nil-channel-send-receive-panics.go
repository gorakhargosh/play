package main

func main() {
	var ch chan int

	// Both will deadlock.
	// fmt.Println(<-ch)
	ch <- 1
}
