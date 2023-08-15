package main

import "fmt"

func main() {
	fmt.Println("Closures")
	fmt.Println("")

	counter := 0
	name := "Budi"

	increment := func() {
		counter++         //mengubah value dari var counter diatas
		name := "Anto"    //kalau dideklarasi ulang, maka akan buat variable baru
		fmt.Println(name) //Anto
	}

	fmt.Println(counter) //0
	increment()
	fmt.Println(counter) //1
	fmt.Println(name)    //Budi

}
