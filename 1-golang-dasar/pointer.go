package main

import "fmt"

/*
	Pointer adalah kemampuan membuat reference ke lokasi data di memory
	yang sama tanpa menduplikasi data yang sudah ada
*/

type Adrress struct {
	City, Province, Country string
}

func main() {
	// address1 := Adrress{"Bandung", "Jawa Barat", "Indonesia"}
	// address2 := &address1 // pointer

	var address1 Adrress = Adrress{"Bandung", "Jawa Barat", "Indonesia"}
	var address2 *Adrress = &address1

	address2.City = "Jakarta"
	fmt.Println(address1)
	fmt.Println(address2)

	// new pointer
	// address2 = &Adrress{"Jakarta", "DKI Jakarta", "Indonesia"}

	// change all value pointer reference
	*address2 = Adrress{"Surabaya", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
