package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation error")
	NotFoundError   = errors.New("Not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "rida" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("rid")

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("ini validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Ini not found error")
		} else {
			fmt.Println("Unknown error")
		}
	}
}
