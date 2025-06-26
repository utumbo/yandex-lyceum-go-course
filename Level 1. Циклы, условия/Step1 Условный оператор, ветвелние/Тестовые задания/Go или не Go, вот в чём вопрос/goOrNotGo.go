package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)

	if input == "Go" {
		fmt.Println("Go!")
	} else {
		fmt.Println("Я знаю только Go!")
	}
}

//Напишите программу, которая принимает на вход строку и,
// если это строка Go, выводит Go!,
// а если нет, — Я знаю только Go!.
