package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// struct method
func (customer Customer) sayHello2() {
	fmt.Println("Hallo ini", customer.Name)

}

func main() {
	//ini komentar single line
	fmt.Println("Struct")
	fmt.Println("Struct itu adalah template data")
	fmt.Println("kalau di java, pengganti class")
	fmt.Println("Struct tidak bisa langsung dipakai")
	//cara 1
	var abdi Customer
	abdi.Name = "Abdillah"
	abdi.Age = 19
	abdi.Address = "Village"
	fmt.Println(abdi) //vilage
	// fmt.Println(abdi.Address)//vilage

	//cara 2
	joko := Customer{
		Name:    "Joko",
		Address: "Bandung",
		Age:     20,
	}
	fmt.Println(joko)

	//cara 3 //tidak disarankan
	budi := Customer{"Budi", "Jakarta", 21} //ini bisa juga, tapi rawan error
	fmt.Println(budi)

	///
	fmt.Println("--------")
	fmt.Println("--------")
	fmt.Println("--------")
	abdi.sayHello2()//ini pemakaian struct func

}
