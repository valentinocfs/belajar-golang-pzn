package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJSON(data interface{}) {
	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}

func TestEncodeJson(t *testing.T) {
	logJSON("John Doe")
	logJSON(123)
	logJSON(true)
	logJSON([]string{"John", "Doe"})
	logJSON(map[string]interface{}{
		"name": "John Doe",
		"age":  30,
	})
}
