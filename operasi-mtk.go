package main

import "fmt"

func main() {
	fmt.Println("Operasi Matematika")
	fmt.Println("+ - * / %")
	var a = 10
	var b = 5
	var ab = a * b
	fmt.Println(ab)

	fmt.Println("")
	fmt.Println("Augmented Assigments")
	fmt.Println("var i = 10")
	fmt.Println("i += 10")
	var i = 10
	fmt.Println(i)	
	i += 10 //jadi 20
	fmt.Println(i)

	fmt.Println("")
	fmt.Println("Unary Operator")
	fmt.Println("++ => a=a+1")
	fmt.Println("-- => a=a-1")
	fmt.Println("- => negative")
	fmt.Println("+ => positif")
	fmt.Println("! => Boolean kebalikan")	
}
