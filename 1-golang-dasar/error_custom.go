package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (nf *notFoundError) Error() string {
	return nf.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "1" {
		return &notFoundError{"data not found"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)

	if err != nil {
		if validationError, ok := err.(*validationError); ok {
			fmt.Println("Validation Error:", validationError.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("Not Found Error:", notFoundError.Error())
		} else {
			fmt.Println("Unknown Error:", err.Error())
		}

		// switch finalError := err.(type) {
		// case *validationError:
		// 	fmt.Println("Validation Error:", finalError.Error())
		// case *notFoundError:
		// 	fmt.Println("Not Found Error:", finalError.Error())
		// default:
		// 	fmt.Println("Unknown Error:", err.Error())
		// }
	} else {
		fmt.Println("Success")
	}
}
