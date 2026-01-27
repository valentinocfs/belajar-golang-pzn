package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "John"
	names[1] = "Doe"
	names[2] = "Wick"

	fmt.Println(names)

	var values = [3]int{
		99,
		98,
		97,
	}
	fmt.Println(values)

	var members = [...]string{
		"John",
		"Jane",
	}
	fmt.Println(len(members))

	//** Tidak ada cara menghapus array di golang
}
