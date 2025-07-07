package main

import "unicode"

func checkPassword(password string) bool {
	if  hasMinimumLength(password) && hasUpper(password){
		return true
	}
	return false
}

func hasMinimumLength(password string) bool {
	pass := len(password)
	if pass >= 8 {

		return true
	}
	return false

}

func hasUpper(password string) bool {
	for _, r := range password {

		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
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
