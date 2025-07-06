package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const queueSize = 5

func main() {
	queue := make([]string, queueSize)
	for i := range queue {
		queue[i] = "-"
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		if strings.ToLower(input) == "конец" {
			printQueue(queue)
			return
		}

		if strings.ToLower(input) == "очередь" {
			printQueue(queue)
			continue
		}

		if strings.ToLower(input) == "количество" {
			printCount(queue)
			continue
		}

		parts := strings.Fields(input)
		if len(parts) < 2 {
			fmt.Printf("Запись на место номер невозможна: некорректный ввод\n")
			continue
		}

		name := strings.Join(parts[:len(parts)-1], " ")
		positionStr := parts[len(parts)-1]

		position, err := strconv.Atoi(positionStr)
		if err != nil {
			fmt.Printf("Запись на место номер %s невозможна: некорректный ввод\n", positionStr)
			continue
		}

		if position < 1 || position > queueSize {
			fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", position)
			continue
		}

		if queue[position-1] != "-" {
			if isQueueFull(queue) {
				fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", position)
			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", position)
			}
			continue
		}

		queue[position-1] = name
	}
}

func printQueue(queue []string) {
	for i, name := range queue {
		fmt.Printf("%d. %s\n", i+1, name)
	}
}

func printCount(queue []string) {
	free := 0
	occupied := 0
	for _, name := range queue {
		if name == "-" {
			free++
		} else {
			occupied++
		}
	}
	fmt.Printf("Осталось свободных мест: %d\n", free)
	fmt.Printf("Всего человек в очереди: %d\n", occupied)
}

func isQueueFull(queue []string) bool {
	for _, name := range queue {
		if name == "-" {
			return false
		}
	}
	return true
}
