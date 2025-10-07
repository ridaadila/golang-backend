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

func (v *notFoundError) Error() string {
	return v.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation error"}
	}

	if id != "rida" {
		return &notFoundError{"Not found error"}
	}

	return nil
}

func main() {
	err := SaveData("rida", nil)

	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("Validation Error: ", validationErr.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("Not Found Error: ", notFoundError.Error())
		} else {
			fmt.Println("Unknown error: ", err.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
