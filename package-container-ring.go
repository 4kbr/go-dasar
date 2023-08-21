package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Package Container/Ring")
	fmt.Println("Package ini adalah implementasi struktur data circular list")
	fmt.Println("Circular list adalah struktur data ring, dimana diakhir element akan kembali kehalaman awal (HEAD)")

	var data *ring.Ring = ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.Itoa(i)
		data = data.Next()
	}
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
	fmt.Println(*data)

	//lengkapnya di golang.org/pkg/container/ring

}
