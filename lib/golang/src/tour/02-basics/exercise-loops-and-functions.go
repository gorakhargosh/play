package main

import (
	"fmt"
	"math"
)

func Sqrt(x, prec float64) float64 {
	z := float64(1)
	for zz := guess(x, z); math.Abs(zz-z) > prec; {
		z = zz
		zz = guess(x, z)
	}
	return z
}

func guess(x, z float64) float64 {
	return z - ((z*z - x) / 2 * z)
}

func main() {
	fmt.Println(Sqrt(2, 0.0004), math.Sqrt(2))
}
