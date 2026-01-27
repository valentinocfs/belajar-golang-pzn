package main

import (
	"fmt"
)

func main() {
	defer endApp()
	runApp(true)
}

func endApp() {
	fmt.Println("End App")
	message := recover()
	if message != nil {
		fmt.Println("Error:", message)
	}
}

func runApp(error bool) {
	if error {
		panic("Application Error")
	}
}

/*
Multine Comment
*/

// Single line comment
