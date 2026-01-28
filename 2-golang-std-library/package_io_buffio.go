package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("Hello, World!\nCall me in the morning")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello, World!")
	_, _ = writer.WriteString("Call me in the morning")
	writer.Flush()
}
