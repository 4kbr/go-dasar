package main

import (
	"fmt"
)

type Address struct {
	City     string
	Province string
	Country  string
}

// pointer di function
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	fmt.Println("Pointer")
	fmt.Println("Secara default golang semua variable menggunakan passing by value")
	fmt.Println("bukan by reference")
	fmt.Println("Artinya datanya diduplicate dari data awal")

	addres1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	addres2 := addres1 //ini tidak pakai pointer

	addres2.City = "Bandung" //data City di address tidak berubah

	fmt.Println(addres2)
	fmt.Println(addres1) //City nya tetap Subang

	fmt.Println("")
	fmt.Println("Pointer adalah kemampuan membuat reference ke lokasi data dimemory yang sama tanpa menduplikasi data yang sudah")
	fmt.Println("Pointer itu pass by refrence")

	addres3 := &addres1 //ini dengan pointer
	addres4 := &addres1 //ini dengan pointer
	addres3.City = "Bekasi"
	addres1.Province = "Bengali"
	fmt.Println(addres1) //data City di addres1 juga berubah
	fmt.Println(addres3)

	fmt.Println("---") //data City di addres1 juga berubah
	fmt.Println("---") //data City di addres1 juga berubah
	// addres3 = &Address{"Bekasi", "Bandung", "Indo"} //cara yang salah untuk merubah semua valuenya sekaligus
	*addres3 = Address{"Jakarta", "DKI Jakarta", "IndoNoesi"} //cara yang benar untuk merubah semua valuenya sekaligus
	fmt.Println(addres1)                                      //ikut berubah juga
	fmt.Println(addres3)
	fmt.Println(addres4) //addres 4 juga ikut berubah

	fmt.Println("")
	fmt.Println("Penggunaan new")
	addres5 := new(Address)
	fmt.Println(addres5)
	addres5.City = "Jakarta"
	addres5.Country = "People"
	addres5.Province = "Asli"
	fmt.Println(addres5)
	ChangeCountryToIndonesia(addres5)
	fmt.Println(addres5)

}
