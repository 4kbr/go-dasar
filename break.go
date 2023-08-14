package main

import "fmt"

func main() {
	fmt.Println("Break and Continue")
	for i := 0; i < 10; i++ {
		fmt.Println("ini print ke", i)
		fmt.Println("SEBELUM DI BREAK")
		if i == 5 {
			break
		}
	}
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ini ganjil", i)
	}

}
