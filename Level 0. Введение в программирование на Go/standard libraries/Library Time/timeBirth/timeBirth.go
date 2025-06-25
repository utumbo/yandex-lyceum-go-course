package main

import (
	"fmt"
	"time"
)

func main(){

	now := time.Now() // текущее время
	timeOfBirth := time.Date(2003, 11, 11, 9, 45, 30, 0, time.UTC)	// время рождения

	diff := now.Sub(timeOfBirth) // разница между текущим временем и рождением
	
	fmt.Println(diff.Hours())	// вывод разницы в часах
}