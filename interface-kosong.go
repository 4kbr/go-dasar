package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1 //return int
	} else if i == 2 {
		return "2" //return string
	} else {
		return false //return boolean
	}
}

func main() {
	//ini komentar single line
	fmt.Println("Interface Kosong")
	fmt.Println("sederhananya kalau function mau return any / dinamis type, pakai interface kosong")
	fmt.Println("")
	fmt.Println(Ups(1))
	fmt.Println(Ups(2))
	fmt.Println(Ups(3))

	// var data int = Ups(1) //ini error
	var data interface{} = Ups(1) //ini success

	fmt.Println(data)

}
