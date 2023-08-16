package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	//ini komentar single line
	fmt.Println("Nil")
	fmt.Println("data null")
	fmt.Println("hanya bisa digunakan dibeberapa tipedata: interface, function, map, slice, pointer dan channel")
	// var person map[string]string = nil
	// fmt.Println(person)

	newPerson := NewMap("Abdi")
	newPersonNill := NewMap("")

	fmt.Println(newPerson)
	fmt.Println(newPersonNill)

	var person = NewMap("Ada")
	if person == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person)
	}

}
