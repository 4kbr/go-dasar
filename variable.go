package main

import "fmt"

func main() {
	fmt.Println("Variable")

	var name string

	name = "Ini variable name"

	fmt.Println(name)

	fmt.Println()

	name = "Ini setelah diganti"
	fmt.Println(name)

	var ini8 int8

	ini8 = 121
	fmt.Println(ini8)

	var frienName = "agung"
	fmt.Println(frienName)

	var age = 19 //auto jadi int
	fmt.Println(age)

	myName := "Bambu"
	fmt.Println(myName)

	myName = "Bukan Bambu"
	fmt.Println(myName)

	var (
		myFullName = "Nama full"
		myAge      = 20
	)
	fmt.Println(myFullName,myAge)

}
