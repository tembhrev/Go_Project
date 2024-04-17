package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100, 0)
	if err != nil {
		log.Println(err)
	}
	log.Println("result of division is", result)
}

func divide(a, b float32) (float32, error) {
	var result float32
	if b == 0 {
		return result, errors.New("division by zero")
	}
	result = a / b
	return result, nil
}
