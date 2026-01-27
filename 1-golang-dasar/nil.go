package main

import "fmt"

func main() {
	map1 := NewMap("")
	if map1 == nil {
		fmt.Println("Map is nil")
	} else {
		fmt.Println(map1)
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}
