package main

import (
	"fmt"
	"strings"
)

// https://pkg.go.dev/strings

func main() {
	str := "Hello World"
	fmt.Println(strings.Contains(str, "Hello"))
	fmt.Println(strings.Count(str, "l"))
	fmt.Println(strings.Replace(str, "World", "Go", 1))
	fmt.Println(strings.Split(str, " "))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.TrimSpace(str))
	fmt.Println(strings.TrimPrefix(str, "Hello"))
	fmt.Println(strings.TrimSuffix(str, "World"))
	fmt.Println(strings.Repeat(str, 2))
	fmt.Println(strings.Join([]string{"Hello", "World"}, " "))
	fmt.Println(strings.Index(str, "World"))
	fmt.Println(strings.ReplaceAll(str, "World", "Go"))
	fmt.Println(strings.Repeat("Hello ", 2))
}
