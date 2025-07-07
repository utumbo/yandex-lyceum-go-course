package main

import (
	"fmt"
	"unicode"
)

func checkPassword(password string) bool {
	if hasMinimumLength(password, 8) && hasUpper(password) {
		return true
	}
	return false
}

func hasMinimumLength(password string, minLength int) bool {
	return len(password) >= minLength
}

func hasUpper(password string) bool {
	for _, r := range password {

		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(hasMinimumLength("a", 10))
	fmt.Println(hasMinimumLength("", 0))
	fmt.Println(hasMinimumLength("aaa", 3))
	fmt.Println(hasMinimumLength("a", 1))
	fmt.Println(hasMinimumLength("afaghlonfwdfbomwdfbj", 10))

	fmt.Println(hasUpper("A"))
	fmt.Println(hasUpper(""))
	fmt.Println(hasUpper("f"))
	fmt.Println(hasUpper("wjefbhqjvoqnihbejbhefhqwlbhdw"))
	fmt.Println(hasUpper("AAAAAAAAAAAAAA"))
	fmt.Println(hasUpper("Aerfbjebrfh"))
	fmt.Println(hasUpper("gerfbjebrfhF"))
	fmt.Println(hasUpper("gerfbJebrfhl"))
}

/*
апишите программу, которая проверяет пароль на соответствие условиям.
Для этого реализуйте функцию checkPassword,
которая принимает на вход строку с паролем и возвращает true,
если пароль им соответствует, и false, если нет.

Для проверки условий реализуйте дополнительную функцию hasMinimumLength,
которая принимает на вход строку с паролем, минимальную длину пароля и проверяет,
длиннее ли пароль необходимого значения. В функции hasMinimumLength эту проверку нужно
выполнить со значением 8.

Также реализуйте функцию hasUpper, которая проверяет,
есть ли в пароле хотя бы одна заглавная буква. Гарантируется,
что на вход будут поданы только латинские буквы.

Все вспомогательные функции также возвращают true,
если пароль соответствует условию, и false, если нет.
*/
