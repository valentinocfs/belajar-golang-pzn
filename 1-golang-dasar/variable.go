package main

import "fmt"

func main() {
	var name string = "John Doe"
	fmt.Println(name)

	// deklarasi awal
	age := 25
	fmt.Println(age)

	// deklarasi ulang
	age = 26
	fmt.Println(age)

	// deklarasi multiple
	var name2, age2 = "John Doe", 25
	fmt.Println(name2, age2)

	// deklarasi multiple
	var (
		name3 = "John Doe"
		age3  = 25
	)
	fmt.Println(name3, age3)

	// constant
	const pi = 3.14
	fmt.Println(pi)
}
