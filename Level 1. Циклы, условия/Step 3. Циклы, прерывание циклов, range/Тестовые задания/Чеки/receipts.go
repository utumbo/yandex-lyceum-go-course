package main

import "fmt"

func main() {
	var products, discount int
	var price, result float64

	fmt.Scan(&products)
	fmt.Scan(&discount)

	for i := 0; i < products; i++ {
		fmt.Scan(&price)
		result += price * float64(100-discount) / 100
	}
	fmt.Print(result)
}

// напишите программу, которой подаётся на вход количество продуктов,
// скидка по карте лояльности (скидка распространяется на все продукты) и цена каждого продукта,
// а программа выводит итоговую стоимость покупок в чеке с учётом скидок.
