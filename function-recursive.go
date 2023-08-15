package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println("Recursive Function")

	fmt.Println("")
	// fmt.Println(factorialLoop(-2)) //error
	fmt.Println(factorialLoop(15)) //ini menggunakan looping
	// fmt.Println(1 * 2 * 3 * 4 * 5)
	fmt.Println(factorialRecursive(15)) //ini menggunakan recursive

}
