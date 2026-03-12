package main

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](b Bag[T]) {
	for _, value := range b {
		fmt.Println(value)
	}
}

func TestBag(t *testing.T) {
	bagString := Bag[string]{"a", "b", "c"}
	PrintBag(bagString)

	bagInt := Bag[int]{1, 2, 3}
	PrintBag(bagInt)
}
