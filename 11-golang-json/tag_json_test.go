package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "Product 1",
		ImageURL: "https://example.com/product-1.jpg",
	}

	jsonString, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonString))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{
		"id": "P0001",
		"name": "Product 1",
		"image_url": "https://example.com/product-1.jpg"
	}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	err := json.Unmarshal(jsonBytes, &product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.ImageURL)
}
