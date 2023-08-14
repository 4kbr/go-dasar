package main

import "fmt"

func main() {
	fmt.Println("Operasi boolean")
	fmt.Println("")
	fmt.Println("Operasi &&")
	fmt.Println("kalau salah satunya false, maka hasilnya false")
	fmt.Println("kalau 22nya true, maka hasilnya true")
	fmt.Println("")
	fmt.Println("Operasi ||")
	fmt.Println("kalau salah satunya true, maka hasilnya true")
	fmt.Println("kalau 22nya false, maka hasilnya false")
	fmt.Println("kalau 22nya false, maka hasilnya false")
	fmt.Println("")
	fmt.Println("Operasi !")
	fmt.Println("operasi kebalikan")
	fmt.Println("!true artinya false")
	fmt.Println("!false artinya true")

	var nilai = 86
	var absensi = 79
	var kkm = 80
	var lulusAbsensi = 80

	fmt.Println("")
	fmt.Println(nilai < kkm && absensi > lulusAbsensi)

}
