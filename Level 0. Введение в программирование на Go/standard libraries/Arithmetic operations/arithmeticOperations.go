package main

import "fmt"

func main() {
	var a = 15
	a += 2         // Эквивалентно a = a + 2
	fmt.Println(a) // 17

	a -= 6         // Эквивалентно a = a - 6
	fmt.Println(a) // 11

	a *= 2         // Эквивалентно a = a * 2
	fmt.Println(a) // 22

	a /= 11        // Эквивалентно a = a / 11
	fmt.Println(a) // 2

	a %= 3         // Эквивалентно a = a % 3
	fmt.Println(a) // 2
}
