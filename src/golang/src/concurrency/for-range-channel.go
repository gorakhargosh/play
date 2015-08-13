package main

import "fmt"

// show A OMIT
type Paratha string

func (p Paratha) String() string { return string(p) + " paratha" }

func ParatheWaaliGali() <-chan Paratha { // HL
	ch := make(chan Paratha)
	go func() {
		ch <- Paratha("Aloo")
		ch <- Paratha("Paneer")
		ch <- Paratha("Amritsari")
		ch <- Paratha("Cukdukoo")
		close(ch) // indicate end of communication // HL
	}()
	return ch
}

func main() {
	for paratha := range ParatheWaaliGali() { // HL
		fmt.Printf("%v bhi khaaya.\n", paratha)
	}
	fmt.Println("\nYaar, pet badh gaya.")
}

// end show A OMIT
