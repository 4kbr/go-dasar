package main

import "fmt"

var endApp = func() {
	fmt.Println("Aplikasi Selesai")
	message := recover() //pemakaian recover yang benar
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
}

//pemakaian recover yang salah
// var runApp = func(error bool) {
// 	defer endApp()
// 	if error {
// 		panic("Aplikasi ERROR")
// 	}
// 	message := recover()

// 	fmt.Println("Error dengan message", message)
// 	fmt.Println("Aplikasi Berjalan")
// }

var runApp = func(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	fmt.Println("Panic function")
	fmt.Println("")
	fmt.Println("Panic function adalah function yang bisa kita gunakan untuk menghentikan program")
	fmt.Println("Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan")
	fmt.Println("Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekus")
	fmt.Println("")
	fmt.Println("-------")
	runApp(false) //ini tidak error
	fmt.Println("-------")
	runApp(true) //error
	fmt.Println("-------")
	fmt.Println("")
	fmt.Println("Recover")
	fmt.Println("Recover adalah function yang bisa kita gunakan untuk menangkap data panic")
	fmt.Println("Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan")

}
