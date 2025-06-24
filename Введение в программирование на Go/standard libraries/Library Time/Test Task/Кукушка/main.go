package main

import (
	"fmt"
	"time"
)

func main() {
	var timeInput string

	fmt.Scan(&timeInput)

	//Парсим время из строки
	t, _ := time.Parse("2006-01-02/15:04:05", timeInput)

	hours := t.Hour()
	minutes := t.Minute()

	fmt.Printf("Текущее время %d часов, %d минут. Ты точно не забыл про важные дела на сегодня?", hours, minutes)
}

//с временем как то тяжело, нужно поглубже изучить данную тему
