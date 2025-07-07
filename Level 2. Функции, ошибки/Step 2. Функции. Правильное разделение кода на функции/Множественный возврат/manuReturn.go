package main

import (
	"fmt"
	"math"
)

func divide(a, b int) (int, int) {
	integer := a / b

	reminder := a % b

	return integer, reminder
}

func calculateAreaAndPerimeter(lenght, wight float64) (float64, float64) {
	area := lenght * wight

	perimeter := 2 * (lenght + wight)

	return area, perimeter
}

func calculateDiagonal(length, width float64) float64 {
	return math.Sqrt(math.Pow(length, 2) + math.Pow(width, 2))
}

func main() {

	integer, reminder := divide(10, 15)
	fmt.Printf("Цлая часть при делении = %d\nОстаток от деления = %d\n", integer, reminder)

	area, perimeter := calculateAreaAndPerimeter(10, 11)

	fmt.Printf("Площадь = %.f\nПериметр = %.f\n", area, perimeter)

	fmt.Println(calculateAreaAndPerimeter(10.0, 5.0))

}
