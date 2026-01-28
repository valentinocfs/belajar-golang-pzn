package main

import (
	"container/list"
	"fmt"
)

// https://pkg.go.dev/container/list

func main() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	l.PushBack(7)
	l.PushBack(8)
	l.PushBack(9)
	l.PushBack(10)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
