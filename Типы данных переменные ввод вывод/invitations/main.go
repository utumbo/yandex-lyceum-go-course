package main

import "fmt"

func main() {
	var name string
	var numberDor, time int
	var password float64

	fmt.Scanln(&name)
	fmt.Scanln(&numberDor)
	fmt.Scanln(&time)
	fmt.Scanln(&password)

	fmt.Printf("Привет, %s! Приглашаю тебя на соревнование по программированию, которое пройдёт, как всегда, в квартире %d. Оно будет длиться примерно %d часа. Не забудь секретный пароль для входа: %d", name, numberDor, password, time)
}
