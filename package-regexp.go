package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Package regexp")
	fmt.Println("Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression")
	fmt.Println("Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2")

	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])u")

	fmt.Println(regex.MatchString("Akulang")) //false
	fmt.Println(regex.MatchString("akulanu")) //true
	fmt.Println(regex.MatchString("aau"))     //true

	search := regex.FindAllString("aku aKiiou Aku amu", 2)
	fmt.Println(search)

	//github github.com/google/re2/wiki/Syntax
	//lengkapnya di golang.org/pkg/regexp
}
