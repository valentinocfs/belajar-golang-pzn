package main

import "fmt"

func main() {
	var data interface{} = "John Doe"
	var data2 interface{} = 100

	// type assertion
	dataString := data.(string)
	fmt.Println(dataString)

	dataInt := data2.(int)
	fmt.Println(dataInt)

	checkType(data)
	checkType(data2)
}

func checkType(data interface{}) {
	switch value := data.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
