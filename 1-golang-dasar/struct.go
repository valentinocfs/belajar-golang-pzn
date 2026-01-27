package main

import "fmt"

type Person struct {
	Name, Address string
	Age           int
}

// struct method
func (person Person) SayHello() {
	fmt.Println("Hello, my name is", person.Name)
}

func main() {
	person := Person{
		Name:    "John Doe",
		Address: "123 Main St",
		Age:     30,
	}
	fmt.Println(person)
	person.SayHello()
}
