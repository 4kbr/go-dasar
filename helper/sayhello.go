package helper

import "fmt"

var version = "ini tidak bisa diakses"
var Application = "ini bisa diakses"
 
func sayHolla(name string) { //ini tidak bisa diimport
	fmt.Println("Hello", name)
}

func SayHolla(name string) { //ini bisa
	fmt.Println("Hello", name)
}

// go env -w GO111MODULE=off
