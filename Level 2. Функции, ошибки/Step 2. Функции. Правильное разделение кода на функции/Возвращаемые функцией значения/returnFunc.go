package main

import (
	"fmt"
	"math"
)

func findDiscriminant(a, b, c float64) float64 {
	discriminant := math.Pow(b, 2) - 4*a*c
	return discriminant
}

func main() {
	fmt.Println(findDiscriminant(1, 4.0, 2))
}
