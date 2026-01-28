package main

import (
	"container/ring"
	"fmt"
)

// https://pkg.go.dev/container/ring

func main() {
	data := ring.New(10)
	for i := 0; i < data.Len(); i++ {
		data.Value = i
		data = data.Next()
	}

	data.Do(func(p interface{}) {
		fmt.Println(p)
	})

}
