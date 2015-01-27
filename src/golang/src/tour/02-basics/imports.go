package main

// You can also write
// import "fmt"
// import "math"
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
}
