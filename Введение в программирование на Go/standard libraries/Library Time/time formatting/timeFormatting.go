package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now() //считывание текущего времени

	fmt.Println("1. Время в формате RFC3339:", now.Format(time.RFC3339))
	fmt.Println("2. Полная дата и время", now.Format("2006-01-02 15:04:05"))
	fmt.Println("3. Краткая дата:", now.Format("2006-01-02"))
	fmt.Println("Время в 24-часовом формате:", now.Format("15:04:05"))
	fmt.Println("5. Время в 12-часовом формате с AM/PM:", now.Format("03:04 PM"))
	fmt.Println("6. День недели:", now.Format("Monday"))
	fmt.Println("7. Сокращённый месяц:", now.Format("Jan"))
}
