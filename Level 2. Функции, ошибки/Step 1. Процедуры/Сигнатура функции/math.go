package main

import (
	"fmt"
	"math"
)

func mathFunc(a, b int64) (int64, int64, int64, int64) {
	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b

	return sum, sub, mul, div
}

func findDiscriminant(a, b, c float64) {
	discriminant := math.Pow(b, 2) - 4*a*c
	fmt.Printf("Дискриминант равен : %.f\n", discriminant)
}

func findXParabolaVertex(a, b int) {
	xParabola := -b / (2 * a)
	fmt.Printf("X параболы равен : %d\n", xParabola)
}

func findYParabolaVertex(a, b, c float64) {
	yParabola := -(math.Pow(b, 2) - 4*a*c) / (4 * a)
	fmt.Printf("Y параболы равен : %.f\n", yParabola)
}

func main() {
	s, su, mu, di := mathFunc(10, 15)
	fmt.Println(s, su, mu, di)

	findDiscriminant(10, 15, 20)
	findDiscriminant(4, 10, -1)

	findXParabolaVertex(10, 5)
	findYParabolaVertex(10, 5, 6)
}
