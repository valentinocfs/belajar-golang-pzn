package main

import "fmt"

func main() {
	type NoKTP string

	var noKTPAndi NoKTP = "123456789"
	fmt.Println(noKTPAndi)
}
