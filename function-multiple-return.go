package main

import "fmt"

func getFullname(name string) (string, string) {
	return "Jyta", name
}

func main() {
	geo, leo := getFullname("Yuki")
	fmt.Println(geo, leo)

	v1, _ := getFullname("Ogke")//tidak pakai v2nya
	fmt.Println("")
	fmt.Println(v1)
}
