package main

import (
	"fmt"
	"time"
)

// https://pkg.go.dev/time

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Local())

	var utc = time.Date(2026, 1, 28, 10, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := time.Now()
	fmt.Println(value)

	var parsed, _ = time.Parse(formatter, value.Format(formatter))
	fmt.Println(parsed)

	var diff = now.Sub(parsed)
	fmt.Println(diff)

	var duration = time.Duration(10 * time.Second)
	fmt.Println(duration)

	var duration1 = time.Second * 100
	var duration2 = time.Minute * 10
	var duration3 = time.Hour * 1

	fmt.Println(duration1.Seconds())
	fmt.Println(duration2.Minutes())
	fmt.Println(duration3.Hours())
}
