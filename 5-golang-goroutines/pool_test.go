package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "-"
		},
	}

	pool.Put("John")
	pool.Put("F")
	pool.Put("Khannedy")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(500 * time.Millisecond)
			pool.Put(data)
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Finish")
}
