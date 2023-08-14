package main

import "fmt"

func main() {
	fmt.Println("Tipe Data Slice")
	fmt.Println("Slice mirip array, tapi ukuran Slice bisa berubah")
	fmt.Println("Slice dan array selalu terkoneksi, data yang disimpan di Slice itu sebenarnya disimpan di array")
	fmt.Println("")
	fmt.Println("Tipe data Slice memiliki 3 data: pointer, length dan capacity")

	fmt.Println("")
	fmt.Println("Membuat slice dari array")
	var arrays [3]int

	arrays[0] = 1
	arrays[1] = 1
	arrays[2] = 1

	names := [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	fmt.Println(names)
	var slice1 = names[3:8]
	var slice2 = names[:8]
	var slice3 = names[1:]
	var slice4 = names[:]
	fmt.Println("")
	fmt.Println("ini Slice")
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println("")
	fmt.Println("kalau slice dirubah, maka arraynya juga ikut ke rubah")
	fmt.Println("")
	fmt.Println(names[0])
	slice4[0] = "test"
	fmt.Println(names) //jadi test
	fmt.Println("Function yang ada di array")
	fmt.Println("len(slice) //untuk mendapat panjang ")
	fmt.Println("cap(slice) //untuk mendapat capacity")
	fmt.Println("append(slice,data) //membuat slice baru dengan menambah ")
	fmt.Println("make([]TypeData,length,capacity)")

	fmt.Println("")
	fmt.Println("kode program make slice")
	newSlice := make([]string, 2, 5)
	newSlice[0] = "satu"
	newSlice[1] = "dua"
	// newSlice[2] = "tiga"
	// newSlice[3] = "empat"
	// newSlice[4] = "lima"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println("")
	fmt.Println("copySlice")
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	fmt.Println("")
	fmt.Println("Perbedaan array dan slice")
	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
