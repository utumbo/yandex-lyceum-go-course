package main

import "fmt"

func main() {
	var numb1, numb2 int

	fmt.Scanln(&numb1)
	fmt.Scanln(&numb2)

	if numb1 > 0 || numb2 > 0 {
		fmt.Println("Как минимум одна переменная больше нуля")
	}
}
