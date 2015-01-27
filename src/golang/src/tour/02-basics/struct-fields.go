package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	// Struct fields are accessed using a dot.
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
