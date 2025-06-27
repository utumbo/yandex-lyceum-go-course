package main

import "fmt"

func main() {
	var firstPassword, secondPassword string

	fmt.Scan(&firstPassword)
	fmt.Scan(&secondPassword)
	//использую if, а не switch потому что мало данных для ветвления и проще использовать if
	if len(firstPassword) > 8 && len(secondPassword) > 8 {
		fmt.Println("Оба пароля надёжные")
	} else if len(firstPassword) < 8 && len(secondPassword) < 8 {
		fmt.Println("Оба пароля ненадёжные")
	} else if len(firstPassword) > 8 && len(secondPassword) < 8 {
		fmt.Println("Только первый пароль надёжный")
	} else {
		fmt.Println("Только второй пароль надёжный")
	}
}

//Выбор надёжного пароля — непростая задача, он должен быть как минимум из 8 символов.
// Помогите программисту Арсению из двух паролей выбрать надёжный.

//Напишите программу, которая принимает два пароля и выводит:

//Оба пароля надёжные
//Оба пароля ненадёжные
//Только первый пароль надёжный
//Только второй пароль надёжный в зависимости от длины пароля.
