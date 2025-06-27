package main

import "fmt"

func main() {
	var num1, num2, num3 float64
	fmt.Scan(&num1, &num2, &num3)

	if num1 == num2 && num2 == num3 && num3 == num1 {
		fmt.Println("Максимальное равенство")
	} else {
		fmt.Println("Не равны")
	}
}

//Напишите программу, которая сравнивает три числа и
//выводит ‘Максимальное равенство‘, если все числа равны или ‘Не равны‘ в противном случае.
