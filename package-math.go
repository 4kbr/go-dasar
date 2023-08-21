package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Package Math")
	fmt.Println("Package untuk matematika")
	fmt.Println("")
	fmt.Println(math.Round(2.3)) // jadi 2

	fmt.Println(math.Round(2.5)) // jadi 3

	fmt.Println(math.Floor(2.9)) // jadi 2

	fmt.Println(math.Ceil(2.1)) // jadi 3

	fmt.Println(math.Max(10, 20)) // return 20
	fmt.Println(math.Min(10, 20)) // return 10
}
