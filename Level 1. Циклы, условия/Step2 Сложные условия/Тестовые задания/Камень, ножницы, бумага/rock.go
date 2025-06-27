package main

import "fmt"

func main() {
	var firstPlayer, secondPlayer string

	fmt.Scan(&firstPlayer)
	fmt.Scan(&secondPlayer)

	switch {
	case firstPlayer == secondPlayer:
		fmt.Println("Ничья")
	case (firstPlayer == "камень" && secondPlayer == "ножницы") ||
		 (firstPlayer == "ножницы" && secondPlayer == "бумага") ||
		 (firstPlayer == "бумага" && secondPlayer == "камень"):
		fmt.Println("Первый игрок победил")
	default:
		fmt.Println("Второй игрок победил")
	}
}

//Напомним, что в этой игре камень побеждает ножницы, ножницы — бумагу, а бумага — камень.

//Программист Арсений очень любит играть в «Камень, ножницы, бумага» и решил разработать
// программу, которая моментально определяет победителя.

//На вход ей подаются два результата игроков, то есть слова: камень, ножницы или бумага.
// Нужно определить, кто из двух игроков выиграл, и вывести один из результатов:

//Первый игрок победил
//Второй игрок победил
//Ничья
