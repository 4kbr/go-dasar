package main

import "fmt"

func main() {
	fmt.Println("Switch")

	nilai := 19
	switch nilai {
	case 90:
		fmt.Println("ini mantap")
	case 19:
		fmt.Println("Apa ini?")
	default:
		fmt.Println("Ini default")
	}

	fmt.Println("")
	fmt.Println("short statement di switch")

	switch length := len("ini nama ku"); length > 9 {
	case true:
		fmt.Println("ini mantap")
	case false:
		fmt.Println("Nama u=sudah Oke")
	}

	fmt.Println("")
	fmt.Println("switch tanpa expression")
	switch {
	case nilai > 90:
		fmt.Println("ini mantap")
	case nilai > 80:
		fmt.Println("sudah Oke")
	case nilai > 70:
		fmt.Println("Lumayan")
	default:
		fmt.Println("Coba lagi ya, jan lupa")
	}
}
