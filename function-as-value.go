package main

import "fmt"

func getGoodBye(name string) string {
	return "goodby " + name
}

func main() {
	fmt.Println("Function as Value")

	newGetGoodBye := getGoodBye

	fmt.Println("")
	fmt.Println(newGetGoodBye("oke"))

}
