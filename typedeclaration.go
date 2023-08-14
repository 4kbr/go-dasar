package main

import "fmt"

func main() {
	fmt.Println("Type deklarasi")

	type NoKTP string
	type Married bool

	var myNoKtp NoKTP = "321487293892898391"
	var marriedStatus Married = false
	var marriedStatus2 Married = "sadasd";
	fmt.Println(myNoKtp)
	fmt.Println(marriedStatus)
	fmt.Println(marriedStatus2)

}
