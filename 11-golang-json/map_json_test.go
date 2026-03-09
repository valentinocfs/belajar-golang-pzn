package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapJson(t *testing.T) {
	jsonString := `{"id": 1, "name": "Product 1", "Price":1000000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["Price"])
}

func TestMapJsonEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":       1,
		"name":     "Product 1",
		"Price":    1000000000,
		"Discount": 0.1,
	}

	jsonString, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonString))
}
