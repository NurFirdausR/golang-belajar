package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("Notfound error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "nur" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("nur")

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation Error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Error Not found")
		} else {
			fmt.Println("Unknown Error")
		}
	}

	fmt.Println("id ditemukan")
}
