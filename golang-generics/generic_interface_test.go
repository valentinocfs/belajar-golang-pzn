package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type Mydata[T any] struct {
	Value T
}

func (m *Mydata[T]) GetValue() T {
	return m.Value
}

func (m *Mydata[T]) SetValue(value T) {
	m.Value = value
}

func TestInterface(t *testing.T) {
	myDataInt := Mydata[int]{Value: 10}
	result := ChangeValue(&myDataInt, 20)
	assert.Equal(t, 20, result)

	myDataString := Mydata[string]{Value: "Hello"}
	result2 := ChangeValue(&myDataString, "World")
	assert.Equal(t, "World", result2)
}
