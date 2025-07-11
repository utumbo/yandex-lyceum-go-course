package main

import (
	"errors"
	"fmt"
)

var DivisionByZeroError = errors.New("division by zero is not allowed")

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, DivisionByZeroError
	}
	result := float64(a) / float64(b)
	return result, nil
}

func main() {
	fmt.Println(Divide(15, 2))
	fmt.Println(Divide(100, 3))
	fmt.Println(Divide(-227, 19))
	fmt.Println(Divide(0, 15))
	fmt.Println(Divide(1, 1))
	fmt.Println(Divide(1, 0))
}

/*
Напишите функцию Divide(a, b int) (float64, error),
которая получает два целых числа и выводит результат их деления.
Если делитель равен нулю, функция должна вернуть в качестве ответа 0 и
сообщение об ошибке (division by zero is not allowed). Если второе число не равно нулю,
верните в качестве ошибки nil. Ошибку нужно вынести в переменную: DivisionByZeroError.
*/
