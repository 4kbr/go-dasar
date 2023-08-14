package main

import "fmt"

func main() {
	fmt.Println("Is first run go in vscode")
	fmt.Println("Yes itsfwork alhamdulillah")

	// tipedata()
}

func tipedata() {
	fmt.Println("tipe data ada 2, `Integer` `Floating Point`")
	// ini := 1
	in8 := 127
	in16 := 32767
	in32 := 2147483647
	in64 := 9223372036854775807

	fmt.Println("nilai maksimum int8: ", in8)
	fmt.Println("nilai maksimum int16: ", in16)
	fmt.Println("nilai maksimum int32: ", in32)
	fmt.Println("nilai maksimum int64: ", in64)

	fmt.Println("")

	fmt.Println("tipe data uint (unsign integer) minimalnya 0, karena nilai minimalnya digeser ke positif ")
	fmt.Println("contoh:")
	fmt.Println("uint8, nilai minimum: ", 0, " nilai maksimum: ", 255)

	fmt.Println("")
	fmt.Println("Tipe data Floating Point")
	fmt.Println("float32 => nilai minimum: 1.18*10^-38  , nilai maksimum: 3.4*10^38")
	fmt.Println("float64 => nilai minimum: 2.23*10^-308  , nilai maksimum: 1.80*10^308")
	fmt.Println("complex64")
	fmt.Println("complex128")

	fmt.Println("")
	fmt.Println("Alias")
	fmt.Println("nama lain dari typedata")
	fmt.Println("byte = uint8")
	fmt.Println("rune = int32")
	fmt.Println("int  = Minimal int32")
	fmt.Println("uint = Minimal uint32")
}
