package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}
func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}
func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	fmt.Println("Package sort")
	fmt.Println("Package sort adalah package yang berisikan utilitas untuk proses pengurutan")
	fmt.Println("Supaya bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface")

	users := []User{
		{"Eko", 30},
		{"Aku", 35},
		{"Joko", 23},
		{"Andi", 21},
		{"Rudi", 19},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)

	//lebih lengkapnya di golang.org/pkg/sort
}
