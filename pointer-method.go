package main

import (
	"fmt"
)

type Man struct {
	Name string
}

// func (man Man) Married() { //kalau tanpa pointer man.Name nya tidak akan berubah
func (man *Man) Married() { //kalau dengan pointer, man.Name nya berubah
	man.Name = "Mr. " + man.Name
	fmt.Println(man.Name)
}

func main() {
	fmt.Println("Pointer di Method")
	fmt.Println("method yang menempel distruct adalah pass by value")
	fmt.Println("direkomendasikan untuk menggunakan pointer di method struct untuk penghematan memmory")

	aku := Man{"Hyaku"}
	aku.Married()
	fmt.Println(aku) //Mr. Hyaku
}
 