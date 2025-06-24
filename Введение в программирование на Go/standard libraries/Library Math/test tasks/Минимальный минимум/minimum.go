package main

import (
	"fmt"
	"math"
)

func main() {
	var arceniiHeight, goshaHeight, iraHeight float64

	//Считываем ввод
	fmt.Scanln(&arceniiHeight)
	fmt.Scanln(&goshaHeight)
	fmt.Scanln(&iraHeight)

	min := math.Min(arceniiHeight, math.Min(goshaHeight, iraHeight)) //math не принимает 3 переменные только попарные поэтому пришлось положить переменную в переменную
	fmt.Println(min)
}
