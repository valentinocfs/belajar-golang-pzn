package main

import (
	"fmt"
	"slices"
)

// https://pkg.go.dev/slices

func main() {
	names := []string{"John", "Jane", "Bob", "Alice", "Tom"}
	values := []int{100, 12, 23, 40, 57}

	fmt.Println(slices.Contains(names, "John"))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Index(names, "Bob"))
}
