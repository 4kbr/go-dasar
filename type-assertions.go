package main

import (
	"fmt"
)

func random() interface{} {
	return "Ups"
}

func main() {
	fmt.Println("Type Assertion")
	fmt.Println("kemampuan merubah tipe data menjadi tipe data yang diinginkan")
	fmt.Println("sering digunakan pada interface kosong")
	fmt.Println("")
	var result interface{} = random()
	// var resultString = result.(string)
	// fmt.Println(resultString)
	switch value := result.(type) { //menggunakan type assertion dengan switch
	case string:
		fmt.Println("value", value, "is string")
	case int:
		fmt.Println("value", value, "is int")
	default:
		fmt.Println("Uknown Type")
	}
}
