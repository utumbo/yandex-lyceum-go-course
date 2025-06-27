package main

import "fmt"

func main() {
	var sign string
	var degrees int

	fmt.Scan(&sign)
	fmt.Scan(&degrees)

	switch {
	case sign == "+" && degrees > 20:
		fmt.Println("Стоит надеть майку и шорты")
	case sign == "+" && degrees >= 10:
		fmt.Println("Стоит надеть штаны и кофту")
	case sign == "+" && degrees == 9:
		fmt.Println("Стоит надеть куртку")
	case sign == "-" && degrees == 3:
		fmt.Println("Стоит надеть куртку")
	case sign == "-" && degrees < -5 || degrees < -5 && degrees < 9 || degrees == 0:
		fmt.Println("Стоит надеть куртку")
	default:
		fmt.Println("Стоит надеть зимнюю куртку")
	}
}

//Напишите программу, которая получает знак + или - и целое число (количество градусов)
// или 0, а выводит:

//Стоит надеть майку и шорты, если градусов больше +20
//Стоит надеть штаны и кофту, если от +10 до +20 градусов
//Стоит надеть куртку, если от -5 до +9
//Стоит надеть зимнюю куртку, если меньше
