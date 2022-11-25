package main

import (
	"errors"
	"fmt"
)

func Connect() error {
	return errors.New("connection failed")
}

func main() {
	// We know Connect can result in an error because it returns an error
	err := Connect()
	if err != nil {
		// handle the error or panic if we want
	}
	numbers := []int{1, 2, 3}
	fmt.Println(numbers[4])
}
