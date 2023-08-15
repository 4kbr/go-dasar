package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging() // ini tetap akan dijalankan diakhir walaupun functionnya error
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("result", result)
	//logging()//ini tidak akan dipanggil kalau function yang sekarang error
}
func main() {
	fmt.Println("Defer function")
	fmt.Println("")
	fmt.Println("adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function function selesai dieksekusi")
	fmt.Println("Defer function akan selalu dieksekusi walaupun terjadi Error difunction yang akan dieksekusi")
	fmt.Println("")
	runApplication(10) //ini bisa
	runApplication(0)  //ini error

}
