package main

// Jika dimulai dari huruf besar maka bisa diakses package lain
// Jika dimulai dari huruf kecil maka tidak bisa diakses package lain
import (
	"1-golang-dasar/helper"
	"fmt"
)

func main() {
	application := helper.Application

	// version := helper.version;
	// goodbye := helper.sayGoodbye("John Doe")

	fmt.Println(application)
}
