package main

import "fmt"

func main() {
	fmt.Println("Tipe Data Array")
	fmt.Println("Array di golang perlu menentukan jumlah data yang bisa ditampung terlebih dahulu")
	fmt.Println("Daya tampung array tidak bisa bertambah setelah array dibuat")
	fmt.Println("Index dimulai dari 0")

	var names [3]string
	names[0] = "oke"
	names[1] = "oke1"
	names[2] = "oke2"
	//names[3] = "oke3" //error
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		1,
		2,
		3,
	}

	fmt.Println(values)
	fmt.Println("")
	fmt.Println("Function Array")
	fmt.Println("len(array) //untuk mendapatkan panjang array")
	fmt.Println("array[index] //mendapatkan data di posisi index")
	fmt.Println("array[index] = value //set value di posisi index")
	// fmt.Println(len([4]int{1,23,4,1}))
	fmt.Println(len(values)) //3

}
