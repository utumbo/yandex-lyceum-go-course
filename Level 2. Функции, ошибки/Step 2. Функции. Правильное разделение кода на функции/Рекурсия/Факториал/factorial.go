package main

import "fmt"

func factorial(f uint) uint {
	if f == 0 {
		return 1
	}

	return f * factorial(f-1)
}

func main() {
	fmt.Println(factorial(5))
}
