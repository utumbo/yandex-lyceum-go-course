package main

import "fmt"

func main(){
	var place1,	place2, place3, place4, place5 string

	for{
		var input string
		fmt.Scan(&input)

		//если пустая строка ввода то пропускаем
		if input == ""{
			continue
		}

		//Выводим все 5 мест очереди
		if input == "очередь"{
			fmt.Println(1, place1)
			fmt.Println(2, place2)
			fmt.Println(3, place3)
			fmt.Println(4, place4)
			fmt.Println(5, place5)
		}


	}
}