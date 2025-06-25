package main

import (
	"fmt"
	"math"
)

func main(){

	var num1, num2 float64

	//считывание двух дробных чисел
	fmt.Scan(&num1)
	fmt.Scan(&num2)

	//программа находит большее и округляет до большего числа
	maxNum := math.Max(num1, num2)

	fmt.Println(math.Floor(maxNum))
}