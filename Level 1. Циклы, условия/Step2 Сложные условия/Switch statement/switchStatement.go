package main

import "fmt"

func main() {
	var weather string
	fmt.Scanln(&weather)

	switch weather {
	case "rain":
		fmt.Println("Погода дождливая")
	case "sunny":
		fmt.Println("Погода солнечная")
	default:
		fmt.Println("Погода не определена")
	}
}
