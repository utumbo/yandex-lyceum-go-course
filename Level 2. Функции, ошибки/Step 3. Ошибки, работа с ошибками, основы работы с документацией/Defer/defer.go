package main

import "fmt"

func main() {
    defer fmt.Println("World") // Отложенный вызов функции вывода в консоль
    fmt.Println("Hello")
}