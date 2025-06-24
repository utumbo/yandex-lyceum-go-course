package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var firstname, lastName, patronymic, dateContract string
	var amount1, amount2, amount3 float64

	fmt.Scan(&dateContract)
	fmt.Scan(&firstname)
	fmt.Scan(&lastName)
	fmt.Scan(&patronymic)
	fmt.Scan(&amount1)
	fmt.Scan(&amount2)
	fmt.Scan(&amount3)

	// Парсинг даты
	t, _ := time.Parse("02.01.2006", dateContract)
	contractDate := t.AddDate(0, 0, 15)
	contractDateStr := contractDate.Format("02.01.2006")

	// Расчет суммы
	total := amount1 + amount2 + amount3
	rubles := math.Floor(total)
	fraction := total - rubles
	pennies := math.Round(fraction * 100)

	overflow := math.Floor(pennies / 100)
	rubles += overflow
	pennies -= overflow * 100

	fmt.Printf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\n", lastName, firstname, patronymic)
	fmt.Printf("Дата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\n", contractDateStr)
	fmt.Printf("Общая сумма выплат составит %.0f руб. %.0f коп.\n\nС уважением,\nГл. бух. Иванов А.Е.", rubles, pennies)
}
