package helper

var version = "1.0.0"
var Application = "Golang"

// tidak bisa diakses di package lain
func sayGoodbye(name string) string {
	return "Goodbye, " + name
}

func SayAnything(name string) string {
	return sayGoodbye(name)
}

func SayHello(name string) string {
	return "Hello, " + name
}
