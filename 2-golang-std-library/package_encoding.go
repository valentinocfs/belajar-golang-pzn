package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// https://pkg.go.dev/encoding

func main() {
	value := "Hello, World!"

	var encoded = base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	var decoded, err = base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(string(decoded))
	}

	csvString := "Name,Age,City\nJohn,30,New York\nJane,25,Los Angeles\nAnon,22,Kentucky"

	var reader = csv.NewReader(strings.NewReader(csvString))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Name", "Age", "City"})
	_ = writer.Write([]string{"John", "30", "New York"})
	_ = writer.Write([]string{"Jane", "25", "Los Angeles"})
	_ = writer.Write([]string{"Anon", "22", "Kentucky"})

	writer.Flush()
}
