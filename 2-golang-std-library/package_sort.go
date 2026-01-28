package main

import (
	"fmt"
	"sort"
)

// https://pkg.go.dev/sort

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (p UserSlice) Len() int {
	return len(p)
}

func (p UserSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p UserSlice) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func main() {
	data := []User{
		{Name: "John", Age: 30},
		{Name: "Jane", Age: 25},
		{Name: "Bob", Age: 35},
	}
	sort.Sort(UserSlice(data))
	fmt.Println(data)
}
