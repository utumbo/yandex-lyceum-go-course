package main

import "fmt"

func main() {
	var number int

	fmt.Scanln(&number)
	//проверяем на квадрат числа сразу в условном операторе
	if square := number * number; square > 50 {
		fmt.Printf("Квадрат числа больше 50 и равен %d", square)
	} else {
		fmt.Printf("Квадрат числа меньше 50 и равен %d", square)
	}
}
