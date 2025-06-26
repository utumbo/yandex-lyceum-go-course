package main

import "fmt"

func main() {
	var x, y int
	fmt.Scanln(&x)
	fmt.Scanln(&y)

	if x > 0 && y > 0 {
		fmt.Println("Обе переменные больше нуля")
	}
}
