package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	fmt.Println("Error Interface")
	fmt.Println("Golang sudah menyediakan handle error dari package errors")

	var contohError error = errors.New("Ups ini error")
	fmt.Println(contohError.Error())

	hasil, err := Pembagi(10, 0) //ini error
	if err == nil {
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("error", err)
	}

}
