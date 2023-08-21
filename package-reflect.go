package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}
type AnotherSample struct {
	Name  string `required:"true" max:"10"`
	Email string
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println("Package reflect")
	fmt.Println("Dalam bahasa pemrograman, biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan")
	fmt.Println("Reflection sangat bergunanketika ingin membuat library yang general sehingga mudah digunakan")

	sample := Sample{"Luigi"}
	var sampleType reflect.Type = reflect.TypeOf(sample)
	fmt.Println(sampleType)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required")) //true
	fmt.Println(sampleType.Field(0).Tag.Get("max"))      //10
	fmt.Println(sampleType.Field(0).Tag.Get("min"))      //kosong

	fmt.Println(IsValid(sample)) //true
	sample.Name = ""
	fmt.Println(IsValid(sample)) //false

	another := AnotherSample{"nama", ""}
	fmt.Println(IsValid(another)) //true
	another = AnotherSample{"", ""}
	fmt.Println(IsValid(another)) //false

	//lengkapnya di golang.org/pkg/reflect
}
