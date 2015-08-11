package main

import (
	"fmt"
	"sync"
)

func raceClosure() { // HL
	var wg sync.WaitGroup
	fmt.Print("racy closure: ")
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() { // HL
			fmt.Print(i) // variable i is shared between 6 goroutines. // HL
			wg.Done()
		}() // HL
	}
	wg.Wait()
	fmt.Println()

}

func noRaceClosurePassArgument() { // HL
	var wg sync.WaitGroup
	fmt.Print("pass as argument: ")
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(n int) { // Use local variable // HL
			fmt.Print(n) // HL
			wg.Done()
		}(i) // pass by value // HL
	}
	wg.Wait()
	fmt.Println()
}

func noRaceClosureUniqueVariable() { // HL
	var wg sync.WaitGroup
	fmt.Print("unique variable: ")
	wg.Add(5)
	for i := 0; i < 5; i++ {
		// Can use closure *carefully* by using a unique variable per iteration.
		// Passing as an argument is much safer and cleaner.
		n := i // HL
		go func() {
			fmt.Print(n) // HL
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println()
}

func main() {
	raceClosure()
	noRaceClosurePassArgument()
	noRaceClosureUniqueVariable()
}
