package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Package strings")
	fmt.Println("package yang banyak berguna untuk tipe data string")
	fmt.Println("Banyak function yang bisa dipakai dari package strings ini")

	fmt.Println("----")
	fmt.Println("")
	fmt.Println(strings.Contains("Aku Bukan Herman", "Aku"))          //true
	fmt.Println(strings.Contains("Aku Bukan Herman", "Budi"))         // false
	fmt.Println(strings.Split("Aku Bukan Herman", " "))               // return []string ["aku", "Bukan","Herman"]
	fmt.Println(strings.ToLower("Aku Bukan Herman"))                  //aku bukan herman
	fmt.Println(strings.ToUpper("Aku Bukan Herman"))                  //AKU BUKAN HERMAN
	fmt.Println(strings.ToTitle("Aku Bukan Herman"))                  //AKU BUKAN HERMAN
	fmt.Println(strings.Trim("     Aku Bukan Herman   a      ", " ")) //Aku Bukan Herman  a

}
