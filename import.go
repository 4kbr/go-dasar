package main

import (
	"fmt"
	"pzn/helper"
)

func main() {
	fmt.Println("Function / variable yang dinamai dengan huruf kecil tidak bisa diimport")
	fmt.Println("")
	helper.SayHolla("Luis")
	fmt.Println(helper.Application)//ini bisa
	// fmt.Println(helper.version)//ini tidak bisa   

 
	// helper.sayHolla("Luis") //ini error
}
