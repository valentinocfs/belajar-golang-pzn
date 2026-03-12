package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	var resultString string = Length[string]("John Doe")
	assert.Equal(t, "John Doe", resultString)

	var resultNumber int = Length[int](100)
	assert.Equal(t, 100, resultNumber)
}

func MultipleParamater[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParamater(t *testing.T) {
	MultipleParamater[string, int]("Windah Basudara", 67)
	MultipleParamater[int, float32](100, 67)
}
