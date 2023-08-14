package main

import (
	"fmt"
)

func main() {
	fmt.Println("If expression")
	fmt.Println("")
	fmt.Println("if kondision {run...}")
	name := "Anto"
	if name == "Anto" {
		fmt.Println("halo Anto")
	} else {
		fmt.Println("halo else")
	}

	nilai := 70
	if nilai >= 80 {
		fmt.Println("Kamu lulus")
	} else {
		fmt.Println("Uji coba lagi")
	}
	fmt.Println("")
	fmt.Println("Else if expression")
	fmt.Println("")
	if nilai >= 85 {
		fmt.Println("Kamu lulus dengan nilai A")
	} else if nilai > 70 {
		fmt.Println("Kamu lulus dengan nilai B")
	} else if nilai == 70 {
		fmt.Println("Kamu nyaris tidak lulus")
	} else {
		fmt.Println("Kamu tidak lulus")
	}

	fmt.Println("")
	fmt.Println("If Short statement")
	if length := 80; length > 9 {
		fmt.Println("Ini true")
	} else {
		fmt.Println("Ini False")
	}
}
