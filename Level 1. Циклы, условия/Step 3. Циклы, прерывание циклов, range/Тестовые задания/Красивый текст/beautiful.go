package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var targetWord string

	fmt.Scan(&N)
	fmt.Scan(&targetWord)

	lowerTargetWord := strings.ToLower(targetWord) //Приводим к нижнему регистру

	count := 0
	for i := 0; i < N; i++ {
		var line string
		fmt.Scan(&line) // считываем всю строку

		words := strings.Fields(line)

		for _, word := range words {
			if strings.ToLower(word) == lowerTargetWord {
				count = count + 1
			}
		}
	}
	fmt.Println(count)

}

/*
1. Читаем число N (сколько будет строк)
2. Читаем слово, которое ищем (targetWord)
3. Приводим targetWord к нижнему регистру
4. В цикле читаем N строк:
5. Разбиваем каждую строку на слова
6. Каждое слово приводим к нижнему регистру
7. Сравниваем с targetWord
8. Если совпадает - увеличиваем счетчик
9. Выводим счетчик
*/

//Помогите ему и напишите программу, которая будет принимать на вход число слов в тексте
// (без знаков препинания и специальных символов), слово, повторения которого нужно подсчитать,
//  и сам текст, а выводить количество повторений.
