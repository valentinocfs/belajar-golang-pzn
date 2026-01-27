package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Bandung", "Jawa Barat", ""}
	fmt.Println(address)
	ChangeCountryToIndonesia(&address)
	fmt.Println(address)
}
