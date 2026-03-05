package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbedString(t *testing.T) {
	fmt.Println("Version:", version)
}

//go:embed draft.png
var logoDraft []byte

func TestEmbedBytes(t *testing.T) {
	err := os.WriteFile("logo-draft.png", logoDraft, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestEmbedMultipleFiles(t *testing.T) {
	a, err := files.ReadFile("files/a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Content A:", string(a))

	b, err := files.ReadFile("files/b.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Content B:", string(b))

	c, err := files.ReadFile("files/c.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Content C:", string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestEmbedPathMatcher(t *testing.T) {
	dir, err := path.ReadDir("files")
	if err != nil {
		panic(err)
	}

	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, err := path.ReadFile("files/" + entry.Name())
			if err != nil {
				panic(err)
			}
			fmt.Println("Content:", string(content))
		}
	}
}
