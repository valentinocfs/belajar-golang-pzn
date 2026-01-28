package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// https://pkg.go.dev/os

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(message)
	if err != nil {
		return err
	}
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 06666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(message)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	createNewFile("test.txt", "Hello, World!\nCall me in the morning")
	content, err := readFile("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(content)
	}
	addToFile("test.txt", "\nThis is new content.")
}
