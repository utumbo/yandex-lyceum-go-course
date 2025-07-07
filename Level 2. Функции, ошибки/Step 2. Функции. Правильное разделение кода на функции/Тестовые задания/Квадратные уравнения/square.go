package main

import (
	"fmt"
	"math"
)

func SquareRoots(a, b, c float64) (float64, float64) {
	if a == 0 && b == 0 {
		return 0, 0
	}

	D := math.Pow(b, 2) - 4*a*c
	if D < 0 {
		return 0, 0
	}

	squarePlus := (-b + math.Sqrt(D)) / (2 * a)
	square := (-b - math.Sqrt(D)) / (2 * a)
	return square, squarePlus
}

func findDiscriminant(a, b, c float64) float64 {
	D := math.Pow(b, 2) - 4*a*c
	return D
}

func main() {
	fmt.Println(SquareRoots(10, 4, 2))
	fmt.Println(findDiscriminant(10, 4, 2))
}

/*
напишите функцию SquareRoots(a, b, c float64) (float64, float64),
которая получает на вход три коэффициента квадратного уравнения и возвращает его
 корни по возрастанию (если корень один, то два одинаковых числа) или два числа 0,
 если корней нет. Для вычисления дискриминанта напишите вспомогательную функцию
 findDiscriminant так, как это сделано на занятии.
*/
