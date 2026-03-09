package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestArrayJson(t *testing.T) {
	customer := Customer{
		FirstName:  "John",
		MiddleName: "Wick",
		LastName:   "Doe",
		Age:        30,
		Married:    true,
		Hobbies:    []string{"Cooking", "Reading", "Gaming"},
	}

	byte, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}

func TestArrayJsonDecode(t *testing.T) {
	jsonString := `
		{
			"FirstName": "John",
			"MiddleName": "Wick",
			"LastName": "Doe",
			"Age": 30,
			"Married": true,
			"Hobbies": ["Cooking", "Reading", "Gaming"]
		}
	`
	jsonBytes := []byte(jsonString)

	customers := &Customer{}
	err := json.Unmarshal(jsonBytes, &customers)
	if err != nil {
		panic(err)
	}

	fmt.Println(customers)
	fmt.Println(customers.Hobbies)
}

func TestArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "John",
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

	byte, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}

func TestArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesia","PostalCode":"88888"}]}`
	jsonBytes := []byte(jsonString)

	customers := &Customer{}
	err := json.Unmarshal(jsonBytes, &customers)
	if err != nil {
		panic(err)
	}

	fmt.Println(customers)
	fmt.Println(customers.FirstName)
	fmt.Println(customers.Addresses)
}

func TestOnlyJsonArray(t *testing.T) {
	jsonString := `[{"Street":"Jalan Lagi Dibangun","Country":"Indonesia","PostalCode":"88888"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesia","PostalCode":"88888"}]`
	jsonBytes := []byte(jsonString)

	addresses := []Address{}
	err := json.Unmarshal(jsonBytes, &addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "St. Hall 123",
			Country:    "Indonesia",
			PostalCode: "9999",
		},
		{
			Street:     "St. Hall 321",
			Country:    "Indonesia",
			PostalCode: "88888",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
