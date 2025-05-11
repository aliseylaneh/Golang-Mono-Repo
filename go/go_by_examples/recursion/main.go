package main

import "fmt"

func fact(value int) int {
	if value == 0 {
		return 1
	}
	return value * fact(value-1)
}
func main() {
	fmt.Println(fact(5))
	var fib func(value int) int
	fib = func(value int) int {
		if value < 2 {
			return value
		}
		return fib(value-1) + fib(value-2)
	}
	fmt.Println(fib(6))
}
