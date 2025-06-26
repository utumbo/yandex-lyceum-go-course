package main

import "fmt"

func main() {
	var rubles, pen, ceil int
	fmt.Scanln(&rubles, &pen, &ceil)

	if rubles+pen > ceil {
		fmt.Println("Сегодня будет вкусный кофе!")
	} else {
		fmt.Println("Стоит подкопить")
	}
}
