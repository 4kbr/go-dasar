package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Package Strconv")
	fmt.Println("Mengubah dari tipe data awal ke tipe data lain")

	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("10000000", 10, 32) //dari string ke int
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err)
	}
	value := strconv.FormatInt(1209990, 16) //dari int ke string
	fmt.Println(value)
	//lebih lengkap di golang.org/pkg/strconv

	intoStr := strconv.Itoa(123021)
	fmt.Println(intoStr)
	strtoInt, err2 := strconv.Atoi("123021")
	if err2 == nil {
		fmt.Println(strtoInt)
	} else {
		fmt.Println(err2)
	}

}
