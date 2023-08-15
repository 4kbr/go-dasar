package main

import "fmt"

type Filter func(string) string

func sayHelloWithFiltered(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" || name == "anjing" {
		return "***"
	}
	return name
}

func main() {
	fmt.Println("Function as Parameter")

	fmt.Println("")
	sayHelloWithFiltered("Akue", spamFilter)
	sayHelloWithFiltered("Anjing", spamFilter)
	fmt.Println("")
	fmt.Println("Buat type Declarations untuk pengganti type func(....)...")

}
