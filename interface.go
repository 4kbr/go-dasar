package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello3(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

// /kalau sudah dibuat seperti ini, maka tidak akan error karena sama sama memiliki GetName
func (person Person) GetName() string {
	return person.Name
}

// /pakai struct yang berbeda juga bisa
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	//ini komentar single line
	fmt.Println("Interface")
	fmt.Println("Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung")
	fmt.Println("Interface berisikan definisi definisi method")
	fmt.Println("biasanya digunakan sebagai kontrak")

	// var eko HasName //ini akan error karena tidak bisa dipakai secara
	// SayHello3(eko)

	// ------ yang benar ------
	var person Person
	person.Name = "Budi"
	SayHello3(person)
	//ini bisa juga

	var animal Animal
	animal.Name = "Kucing"
	SayHello3(animal)

}
