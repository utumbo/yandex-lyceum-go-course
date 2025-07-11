package main

import "errors"

var DivisionByZeroError = errors.New("division by zero is not allowed")

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, DivisionByZeroError
	}
	result := float64(a) / float64(b)
	return result, nil
}
