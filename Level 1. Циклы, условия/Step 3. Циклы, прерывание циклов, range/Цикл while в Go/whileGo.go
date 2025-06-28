// О как, а его нет, место while то же используется for
package main

import "fmt"

func main() {
	i := 1000

	for i > 0 {
		fmt.Printf("Я гуль %d -7\n", i)
		i -= 7
	}
}
