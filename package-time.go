package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Package time")
	fmt.Println("biasa digunakan yang berhubungan dengan waktu")

	var now time.Time = time.Now()
	fmt.Println(now)
	fmt.Println(now.Unix())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2024, 8, 19, 12, 23, 50, 23, time.UTC)
	fmt.Println(utc)

	layout := "2006-01- 02" //layoutnya perlu pake contoh
	parse, _ := time.Parse(layout, "2023-10-02")
	fmt.Println(parse)

}
