package main

import (
	"errors"
	"fmt"
)

func sqrLessThanEqual(arg int) (int, error) {
	if arg > 100 {
		return -1, ErrorHigherThanHundred

	}
	return arg * arg, nil
}

// ErrorHigherThanHundred Sentinel error or custom errors
var ErrorHigherThanHundred = fmt.Errorf("the value is higher than 100")
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrorPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea, %w", ErrorPower)
	}
	return nil
}
func main() {
	value, err := sqrLessThanEqual(1000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
	fmt.Println(makeTea(4))
	fmt.Println(errors.Is(makeTea(4), ErrorPower))
	fmt.Println(errors.Is(makeTea(4), ErrorHigherThanHundred))

	for _, i := range []int{1, 1000, 2200, 2} {
		if v, e := sqrLessThanEqual(i); e != nil {
			fmt.Println("Progress failed, ", e)
		} else {
			fmt.Println("Progress worked, ", v)
		}
	}
	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrorPower) {
				fmt.Println("Now it is dark.")
			} else if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea.")
			} else {
				fmt.Printf("Unknow error, %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready")
	}
}
