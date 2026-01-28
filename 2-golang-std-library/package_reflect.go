package main

import (
	"fmt"
	"reflect"
)

// https://pkg.go.dev/reflect

type Sample struct {
	Name string `required:"true" max:"100"`
}

type Person struct {
	// struct tag
	Name    string `required:"true" max:"100"`
	Address string `required:"true" max:"255"`
	Email   string `required:"true" min:"5"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType)
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println("Field", valueField.Name, "with type", valueField.Type)
		fmt.Println("Tag", valueField.Tag)
	}
}

func isValid(value any) bool {
	result := true
	var valueType reflect.Type = reflect.TypeOf(value)
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)

		// required
		if valueField.Tag.Get("required") == "true" {
			if reflect.ValueOf(value).Field(i).Interface() == "" {
				result = false
				if result == false {
					return result
				}
			}
		}
	}
	return result
}

func main() {
	readField(Sample{Name: "John"})
	readField(Person{Name: "Jane", Address: "123 Main St", Email: "jane@jane.com"})
	fmt.Println(isValid(Sample{Name: "John"}))
	fmt.Println(isValid(Person{Name: "Jane", Address: "123 Main St", Email: "jane@gmail.com"}))
}
