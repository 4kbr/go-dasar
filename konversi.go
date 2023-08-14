package main

import "fmt"

func main() {
	fmt.Println("Konversi tipe data")

	var nilai32 int32 = 200
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai64) //bahaya konver dari int yang lebih besar ke lebih kecil

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	fmt.Println("")
	fmt.Println("konversi String")
	var name = "Aku"
	var a = name[0]
	fmt.Println(a)
	var aString = string(a)
	fmt.Println(aString)
	
}
