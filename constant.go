package main

import "fmt"

func main() {
	fmt.Println("Constant adalah variable yang tidak bisa diubah")
	const firstName string = "First"
	const lastName = "Last"
	const value = 1000

	fmt.Println(firstName, lastName, value)
	fmt.Println("")
	fmt.Println("Deklarasi multiple constant")

	const (
		name = "My Name"
		age  = 23
	)
	fmt.Println(name, age)

}
