package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Golang sudah menyiapkan banyak package bawaan")
	fmt.Println("ada banyak yang bisa dicek di: golang.org/pkg/os")
	fmt.Println("---")
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)
	// fmt.Println(args[1])//perlu argumen
	// cara pakai argument  diterminal : go run package-os.go Aku ges
	fmt.Println("----")
	fmt.Println("Hostname")
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname:", name)
	} else {
		fmt.Println("Errpr:", err)
	}

}
