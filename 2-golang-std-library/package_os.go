package main

import (
	"fmt"
	"os"
)

// https://pkg.go.dev/os

func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("Hostname :", hostname)
	}
}
