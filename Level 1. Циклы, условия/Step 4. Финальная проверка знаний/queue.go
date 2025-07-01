package main

import "fmt"

func main(){
	var place1,	place2, place3, place4, place5 string
	fmt.Scan(place1)
	fmt.Scan(place2)
	fmt.Scan(place3)
	fmt.Scan(place4)
	fmt.Scan(place5)
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
			continue
		}

		if input == "количество"{
			occupied := 0

			if place1 != "" {occupied += 1}
			if place2 != "" {occupied += 1}
			if place3 != "" {occupied += 1}
			if place4 != "" {occupied += 1}
			if place5 != "" {occupied += 1}

			fmt.Printf("Осталось свободных мест: %d", 5-occupied)
			fmt.Printf("Всего человек в очереди: %d", occupied)
			continue
		}

		if input == "конец"{
			fmt.Println(1, place1)
			fmt.Println(2, place2)
			fmt.Println(3, place3)
			fmt.Println(4, place4)
			fmt.Println(5, place5)
			return
		}


	}
}