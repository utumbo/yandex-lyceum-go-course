package main

import (
	"fmt"
	"time"
)

func main() {
	var timeFuture, timePast string

	//Считываем ввод
	fmt.Scan(&timeFuture)
	fmt.Scan(&timePast)

	tF, _ := time.Parse("2006-01-02", timeFuture)
	tP, _ := time.Parse("2006-01-02", timePast)

	res := tF.Year() - tP.Year()

	fmt.Printf("%d year ago", res)
}
