package main

import "fmt"

func main() {
	fmt.Println("For loop")

	counter := 1
	for counter <= 10 {
		fmt.Println("Print ke ", counter)
		counter++
	}

	fmt.Println("")
	fmt.Println("For loop dengan statement")
	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("ini namanya for loop statement ke: ", counter2)
	}

	fmt.Println("")
	fmt.Println("contoh for slice")
	slice := []string{"Aku", "Kamu", "kita", "dia"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("")
	fmt.Println("For Range")
	for i, name := range slice {
		fmt.Println("di index", i, "ada", name)
		// for _, name := range slice {
		// fmt.Println("di index ada", name)
	}

}
