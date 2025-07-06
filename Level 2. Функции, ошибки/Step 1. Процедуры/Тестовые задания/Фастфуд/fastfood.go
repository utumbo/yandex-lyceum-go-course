package main

import "fmt"

func BuyFries(size string){
	var price int

	switch size{
	case "S":
		price = 49
	case "M":
		price = 89
	case "L":
		price = 99
	default:
		fmt.Println("Некорректный размер")
		return
	}
	printPrice(price, "Картошка фри")
}

func BuyCola(size string){
	var price int

	switch size{
	case "S":
		price = 99
	case "M":
		price = 119
	case "L":
		price = 139
	default:
		fmt.Println("Некорректный размер")
		return
	}
	printPrice(price, "Кола")
}

func printPrice(prise int, product string){
	fmt.Printf("%s будет стоить %d рублей\n", product, prise)
}

func main(){
	//Примеры вызовов
	BuyFries("M")
	BuyCola("S")
	BuyFries("L")
	BuyCola("L")
}

/*
Напишите две процедуры: BuyFries() и BuyCola() для покупки соответствующего блюда. 
На вход каждой из процедур подаётся размер в виде строки: S, M или L, 
а вывести нужно Картошка фри будет стоить *количество* рублей или 
Кола будет стоить *количество* рублей.

Таблица цен:

Картошка фри S — 49 рублей
Картошка фри M — 89 рублей
Картошка фри L — 99 рублей

Кола S — 99 рублей
Кола M — 119 рублей
Кола L — 139 рублей
В случае неправильного ввода размера нужно вывести Некорректный размер.

Во избежание повторений также напишите вспомогательную процедуру printPrice(), 
которая получает на вход целое число рублей и название продукта, а выводит нужную фразу.

То есть в процедурах BuyFries() и BuyCola() нужно только подсчитать итоговую сумму, 

а вывести результат с помощью вызова printPrice().
*/