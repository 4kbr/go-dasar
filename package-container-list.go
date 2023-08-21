package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("Package Container/list")
	fmt.Println("Untuk buat linkedlist")
	fmt.Println("---")
	data := list.New()
	data.PushBack("Aku")            //menambah input "Aku" ke paling belakang
	data.PushBack("Kamu")           //menambah input "Kamu" ke paling belakang
	data.PushBack("Kita")           //menambah input "Kita" ke paling belakang
	fmt.Println(data.Front().Value) //Aku
	fmt.Println(data.Back().Value)  //Kita
	data.PushFront("Dia")           //menambah input "Dia ke paling depan"
	fmt.Println(data.Front().Value) //Dia

	///dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println("ini", element.Value) //dari dia ke kita
	}
	fmt.Println("-----")
	///dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println("ini", element.Value)
	}
	//lebih lengkapnya di golang.org/pkg/container/list 
}
