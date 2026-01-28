package main

import (
	"fmt"
	"regexp"
)

// https://pkg.go.dev/regexp

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("^([a-zA-Z0-9._%+-]+)@([a-zA-Z0-9.-]+)\\.([a-zA-Z]{2,})$")
	fmt.Println(regex.MatchString("test@gmail.com"))

	var regex1 *regexp.Regexp = regexp.MustCompile(`J([a-zA-Z]+)n`)
	fmt.Println(regex1.MatchString("John"))
	fmt.Println(regex1.MatchString("Jane"))

	fmt.Print(regex1.FindAllString("John Jane Jimin Jose Jayyan", 5))
}
