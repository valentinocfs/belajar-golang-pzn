package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

// Untuk method direkomendasikan selalu gunakan pointer

func main() {
	man := Man{"John Doe"}
	man.Married()
	fmt.Println(man)
}
