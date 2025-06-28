package main

import "fmt"

func main() {
	var student int
	var scores float64

	fmt.Scan(&student)

	for i := 0; i < student; i++ {

		fmt.Scan(&scores)

		switch {
		case (scores >= 90 && scores <= 100):
			fmt.Println(5)
		case (scores >= 75 && scores <= 89):
			fmt.Println(4)
		case (scores >= 50 && scores <= 74):
			fmt.Println(3)
		case (scores <= 49 && scores >= 0):
			fmt.Println(2)
		default:
			fmt.Println("Неверный балл")
		}
	}

}

// Напишите программу, которой подаётся на вход количество учеников и балл каждого ученика.
// На основании этих данных выведите оценку:

// 5, если балл от 90 до 100
// 4, если балл от 75 до 89
// 3, если балл от 50 до 74
// 2, если балл от 0 до 49
// Неверный балл в иных случаях
