package main

import "fmt"

func main(){
	for i := 0; i < 10; i++{
		for j := 0; j < 10; j++ {
			if i * j == 10 || i * j == 0{ //избавляемся от нулей в начале
				continue
			}
			fmt.Println(i * j)
		}
	}
}