package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, err := os.Open("Customer.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}

func TestStreamEncoder(t *testing.T) {
	file, err := os.Create("CustomerOut.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(file)

	// data
	customer := &Customer{
		FirstName:  "John",
		MiddleName: "Wick",
		LastName:   "Doe",
		Age:        30,
		Married:    true,
		Hobbies:    []string{"Cooking", "Reading", "Gaming"},
		Addresses: []Address{
			{
				Street:     "St. Hall 123",
				Country:    "Indonesia",
				PostalCode: "12345",
			},
			{
				Street:     "St. Hall 321",
				Country:    "Indonesia",
				PostalCode: "54321",
			},
		},
	}

	encoder.Encode(customer)
}
