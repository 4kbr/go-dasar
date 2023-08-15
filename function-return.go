package main

import "fmt"

func getHello(name string) string {
	return "halo " + name
}

func main() {
	geo := getHello("Yuki")
	fmt.Println(geo)
}
