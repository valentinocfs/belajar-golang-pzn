package main

import (
	"fmt"
	"strconv"
)

// https://pkg.go.dev/strconv

func main() {
	str := "1"
	i, _ := strconv.Atoi(str)
	fmt.Println(i)

	boolean, _ := strconv.ParseBool("true")
	fmt.Println(boolean)

	f, _ := strconv.ParseFloat("1.2", 64)
	fmt.Println(f)

	digit, _ := strconv.ParseInt("1", 10, 64)
	fmt.Println(digit)

	u, _ := strconv.ParseUint("1", 10, 64)
	fmt.Println(u)

	fmt.Println(strconv.Itoa(1))
}
