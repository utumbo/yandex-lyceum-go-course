package main

import "fmt"

func main() {
	for i, later := range "Hello, World!" {
		fmt.Println(i, string(later))
	}
}

//Работает только с английским языком
