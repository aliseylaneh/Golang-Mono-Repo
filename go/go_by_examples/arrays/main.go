package main

import "fmt"

func main() {
	// Implementing an array using dynamic sizing of slice with triple dots.
	// triple dots means that goland is going to initialize the array's length with the count of elements within it.
	s := [...]int{2, 34, 1, 3, 41, 55}
	fmt.Printf("Elements: %v, Length: %v, Capacity: %v\n", s, len(s), cap(s))
}
