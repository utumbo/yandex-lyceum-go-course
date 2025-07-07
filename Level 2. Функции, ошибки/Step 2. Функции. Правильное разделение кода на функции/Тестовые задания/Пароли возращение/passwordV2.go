package main

import (
	"fmt"
	"unicode"
)

func ratePassword(password string) string {
	minLanght := 8
	score := 0

	if hasMinimumLength(password, minLanght) {
		score++
	}
	if hasUpper(password) {
		score++
	}
	if hasLowerCase(password) {
		score++
	}

	switch score {
	case 1:
		return "Слабый пароль"
	case 2:
		return "Средний пароль"
	case 3:
		return "Сложный пароль"
	default:
		return "Пароль недостаточно безопасен, придумайте новый"
	}
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

func hasLowerCase(password string) bool {
	for _, r := range password {
		if unicode.IsLower(r) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(ratePassword("password"))
}
