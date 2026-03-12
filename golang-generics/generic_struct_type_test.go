package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello, " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[int]{First: 1, Second: 2}
	fmt.Println(data)

	dataString := Data[string]{First: "Hello", Second: "World"}
	fmt.Println(dataString)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{First: "Windah", Second: "Barusadar"}

	assert.Equal(t, "Windah Barusadar", data.ChangeFirst("Windah Barusadar"))
	assert.Equal(t, "Hello, Windah Barusadar", data.SayHello("Windah Barusadar"))
}
