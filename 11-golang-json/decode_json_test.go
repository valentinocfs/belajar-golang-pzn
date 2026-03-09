package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{
		"FirstName": "John",
		"MiddleName": "Wick",
		"LastName": "Doe",
		"Age": 30,
		"Married": true
	}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
}
