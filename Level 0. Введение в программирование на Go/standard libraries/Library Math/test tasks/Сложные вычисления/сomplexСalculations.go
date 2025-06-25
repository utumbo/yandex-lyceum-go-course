package main

import (
	"fmt"
	"math"
)

func main(){
	//Задаем переменные
	var x, y, z, m, n float64

	//Считываем данные 
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	fmt.Scanln(&z)
	fmt.Scanln(&m)
	fmt.Scanln(&n)

	ressult := ((5 * x) * math.Sin(2 * y) / (z + math.Pow(n, math.Log(m))))// (5x * sin(2y)) / (z + n ^ ln(m))

	fmt.Println(ressult)//Вывод результата
}