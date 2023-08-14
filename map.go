package main

import (
	"fmt"
)

func main() {
	fmt.Println("Tipe Data Map")
	person := map[string]string{
		"nama":   "Akuubar",
		"alamat": "Jakarta",
	}
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])
	fmt.Println(person)
	person["status"] = "stay"
	fmt.Println(person)
	fmt.Println("")
	fmt.Println("Function di Map")
	fmt.Println(len(person))

	book := make(map[string]string)
	book["nama"] = "Aku dan bidadari"
	book["penulis"] = "Amanot"
	fmt.Println(book)
	delete(book, "nama") //mendelete valu di key "nama"
	fmt.Println(book)
}
