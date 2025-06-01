package main

import "fmt"

type argError struct {
	arg     int
	message string
}

func (ae *argError) Error() string {
	return fmt.Sprintf("%d - %s", ae.arg, ae.message)
}

func sqrLessThanEqualCustomError(arg int) (int, error) {
	if arg > 100 {
		return -1, &argError{arg: arg, message: "the value is higher than 100"}
	}
	return arg * arg, nil
}
