package main

import "fmt"

func main() {
	names := []string{"John", "Doe", "Wick", "Doe"}

	slice1 := names[2:4]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	daySlice := days[5:]
	fmt.Println(daySlice)

	daySlice[0] = "New Saturday"
	daySlice[1] = "New Sunday"
	fmt.Println(daySlice)
	fmt.Println(days)

	daySlice2 := append(daySlice, "New Days")
	daySlice2[0] = "Old Saturday"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// using make
	newSlice := make([]string, 2, 5)
	newSlice[0] = "John"
	newSlice[1] = "Doe"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Wick")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Jane"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// copy
	copySlice := make([]string, len(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
