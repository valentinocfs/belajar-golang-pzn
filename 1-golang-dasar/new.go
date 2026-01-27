package main

import "fmt"

type Adrress struct {
	City, Province, Country string
}

func main() {
	// var address *Adrress = &Adrress{}
	// var address2 *Adrress = address

	var address *Adrress = new(Adrress)
	var address2 *Adrress = address

	address2.City = "Bandung"
	fmt.Println(address)
	fmt.Println(address2)
}
