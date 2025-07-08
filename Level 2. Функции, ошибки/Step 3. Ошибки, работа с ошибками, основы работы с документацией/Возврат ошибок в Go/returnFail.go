package main

import (
	"errors"
	"fmt"
)

var (
	errDivision = errors.New("division by zero") // задавать ошибки лучше как константы
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errDivision // возвращаем ошибку
	}
	return a / b, nil
}

func main() {
	fmt.Println(divide(10, 0)) // пример с ошибкой
	fmt.Println(divide(10, 2)) // пример без ошибки
}
