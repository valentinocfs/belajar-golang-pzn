package main

import (
	"flag"
	"fmt"
)

// https://pkg.go.dev/flag

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "", "Put your database password")
	port := flag.Int("port", 3306, "Put your database port")

	flag.Parse()

	// go run package_flag.go -username=root -host=localhost -password=123 -port=543

	fmt.Println("Host:", *host)
	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Port:", *port)
}
