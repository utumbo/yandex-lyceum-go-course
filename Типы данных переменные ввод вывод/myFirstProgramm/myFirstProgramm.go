package main

import"fmt"

func main(){
	var firstName, lastName string
	var age int
	var mark float64

		//Считывание имени
	fmt.Println("Введите имя: ")
	fmt.Scanln(&firstName)

		//Считывание фамилии
	fmt.Println("Введите фамилию: ")
	fmt.Scanln(&lastName)

		//Счяитывание age
	fmt.Println("Введите возраст: ")
	fmt.Scanln(&age)

		//Считывание среднего бала
	fmt.Println("введите средний бал по любимому предмету: ")
	fmt.Scanln(&mark)

	//задаем константу

	const secretNumber = 7126405

	fmt.Println(fmt.Sprintf("Привет, я робот-предсказатель! Вижу, что тебе сейчас %d, а зовут тебя %s %s. Уникальное число твоего предсказания %d. А средняя оценка по твоему любимому предмету %f. Как я догадался? Пусть это останется секретом моего мастерства!", 
	age, firstName, lastName, secretNumber ,mark,
	),
)
}
//нанписать код робота предсказателя который считывает age,name,secretNumber(задать константу),mark(средний бал)