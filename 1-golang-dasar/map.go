package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "John Doe"
	// person["age"] = "30"
	// person["address"] = "123 Main St"

	person := map[string]string{
		"name":    "John Doe",
		"age":     "30",
		"address": "123 Main St",
	}

	fmt.Println(person)
	fmt.Println(person["name"])

	length := len(person)
	fmt.Println(length)

	delete(person, "address")
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "John Doe"
	fmt.Println(book)
}
