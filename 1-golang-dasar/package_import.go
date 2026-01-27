package main

import (
	"1-golang-dasar/helper"
	"fmt"
)

// package adalah sebuah direktori folder

func main() {
	hello := helper.SayHello("John Doe")
	fmt.Println(hello)
}
