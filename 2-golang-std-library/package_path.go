package main

import (
	"fmt"
	"path"
)

// https://pkg.go.dev/path

func main() {
	fmt.Println(path.Join("/a", "/b", "/c"))
	fmt.Println(path.Base("/a/b/c"))
	fmt.Println(path.Dir("/a/b/c"))
	fmt.Println(path.Ext("/a/b/c.txt"))
	fmt.Println(path.IsAbs("/a/b/c"))
	fmt.Println(path.Clean("/a/b/c"))
	fmt.Println(path.Split("/a/b/c"))
	fmt.Println(path.Match("/a/b/c", "/a/b/c"))
	fmt.Println(path.Match("/a/b/c", "/a/b/c.txt"))
}
