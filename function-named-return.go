package main

import "fmt"

func getUser() (username, email string, age int) {
	username = "username1"
	email = "emailuser@as.com"
	age = 14
	return
}

func main() {

	fmt.Println("function named return")
	// username, email, age := getUser() ini bisa
	u, a, b := getUser() //ini bisa juga
	fmt.Println(u)
	fmt.Println(a)
	fmt.Println(b)
}
