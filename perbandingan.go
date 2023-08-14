package main

import "fmt"

func main() {
	fmt.Println("Operasi Perbandingan")
	fmt.Println(">  => Lebih dari")
	fmt.Println("<  => kurang dari")
	fmt.Println(">=  => lebih besar atau sama dengan")
	fmt.Println("<=  => lebih kecil atau sama dengan")
	fmt.Println("==  => sama dengan")
	fmt.Println("!=  => tidak sama dengan")

	var val1 = 100
	var val2 = 150
	var isVal1Smaller = val1<val2 //true
	fmt.Println(isVal1Smaller)
}
