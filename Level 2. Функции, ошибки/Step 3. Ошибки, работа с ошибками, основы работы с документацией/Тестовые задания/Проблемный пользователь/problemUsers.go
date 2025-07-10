package main

import "fmt"

func main() {
	var age int
	var name string
	var passportSeriesAndNumber string

	_, err := fmt.Scanln(&age, &name, &passportSeriesAndNumber)
	if err != nil {
		fmt.Println("Ошибка: некорректный ввод")
		return
	}

	if age < 14 || age > 150 {
		fmt.Println("Ошибка: невалидный возраст")
		return
	}

	if len(name) < 2 {
		fmt.Println("Ошибка: невалидное имя")
		return
	}

	if len(passportSeriesAndNumber) != 10 {
		fmt.Println("Ошибка: невалидная серия и номер паспорта")
		return
	}
	fmt.Println(fmt.Sprintf("Имя: %s. Возраст: %d. Серия и номер паспорта: %s", name, age, passportSeriesAndNumber))
}

/*
Если происходит ошибка в функции Scan, нужно вывести: Ошибка: некорректный ввод.
Если возраст меньше 14 или больше 150, нужно вывести: Ошибка: невалидный возраст
(нижняя граница 14, потому что паспорт в России выдают с 14 лет).
Если имя короче двух символов, нужно вывести: Ошибка: невалидное имя.

Также проверьте, что серия и номер паспорт состоят из 10 знаков.
 В случае ошибки выведите: Ошибка: невалидная серия и номер паспорта.
*/
