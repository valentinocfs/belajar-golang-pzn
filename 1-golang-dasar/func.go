package main

import "fmt"

type Blacklist func(string) bool

func main() {
	sayHello()

	result := getHello("John")
	fmt.Println(result)

	// ignore second return value
	firstName, _ := getFullName()
	fmt.Println(firstName)

	sum := sumAll(1, 2, 3, 4, 5)
	fmt.Println(sum)

	// named return value
	goodbye := getGoodbye("John")
	fmt.Println(goodbye)

	sayHelloWithFilter("anjay", spamFilter)

	// anonymous function (1)
	blacklist := func(name string) bool {
		return name == "anjay"
	}
	registerUser("anjay", blacklist)

	// anonymous function (2)
	registerUser("anjay", func(name string) bool {
		return name == "anjay"
	})

	// recursive function
	fmt.Println(factorial(5))

	// closure
	counter := 0
	increment := func() int {
		counter++
		return counter
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

func sayHello() {
	fmt.Println("Hello")
}

// single return value
func getHello(name string) string {
	return "Hello " + name
}

// multiple return values
func getFullName() (string, string) {
	return "John", "Doe"
}

// variadic parameters
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// named return value
func getGoodbye(name string) string {
	return "Goodbye " + name
}

// function as parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello " + filter(name))
}

func spamFilter(name string) string {
	if name == "anjay" {
		return "..."
	}
	return name
}

// anonymous function
func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are in blacklist")
	} else {
		fmt.Println("Welcome", name)
	}
}

// recursive function
func factorial(x int) int {
	if x == 1 {
		return 1
	}
	return x * factorial(x-1)
}
