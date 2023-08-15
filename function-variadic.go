package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	fmt.Println("function variadic ....string")

	total := sumAll(10, 29, 230, 123, 0, 123, 1, 123)
	fmt.Println("")
	fmt.Println(total)
	slice := []int{1, 2, 4, 6, 123, 3, 123, 53, 12, 531, 14}
	total2 := sumAll(slice...)
	fmt.Println(total2)
}
