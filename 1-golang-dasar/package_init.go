package main

import (
	"1-golang-dasar/database"
	_ "1-golang-dasar/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
