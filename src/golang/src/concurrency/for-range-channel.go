package main

import "fmt"

// show A OMIT
type Paratha string

func ParatheWaaliGali() <-chan Paratha { // HL
	ch := make(chan Paratha)
	go func() {
		ch <- Paratha("Aloo")
		ch <- Paratha("Paneer")
		ch <- Paratha("Amritsari")
		ch <- Paratha("Chana Daal")
		ch <- Paratha("Chicken")
		ch <- Paratha("Garlic")
		close(ch) // indicate end of communication // HL
	}()
	return ch
}

func main() {
	for paratha := range ParatheWaaliGali() { // HL
		fmt.Printf("%s paratha bhi khaaya.\n", paratha)
	}
	fmt.Println("\nYaar, pet badh gaya.")
}

// end show A OMIT
