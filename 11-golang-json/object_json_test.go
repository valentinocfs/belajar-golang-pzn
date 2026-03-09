package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestObjectJson(t *testing.T) {
	customer := Customer{
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
		},
	}

	byte, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}
