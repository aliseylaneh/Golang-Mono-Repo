package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string = []string{"Ali", ""}
	fmt.Printf("len: %v, value: %v\n", len(s), s)
	s[1] = "Ali Seylaneh"
	fmt.Println(s == nil)
	fmt.Println(s)
	//	---------------
	s = make([]string, 5)
	fmt.Println(s)
	fmt.Println(s == nil)
	v := make([]int, 3, 5)
	v[0] = 1
	v[1] = 2
	v[2] = 3
	//v[3] = 4
	v = append(v, 4)
	v = append(v, 5)
	fmt.Println(cap(v), len(v))
	// ----------------
	c := make([]int, len(v))
	copy(c, v) // All the attributes of the slice is copied length, cap and elements.
	fmt.Println(c, cap(c))
	v[2] = 33
	fmt.Println(v)
	fmt.Println(c)
	// ------------
	fmt.Println("Slicing in slice")
	fmt.Println(c[1:3], cap(c))
	fmt.Println(c[:4]) // Excluding the index of 4 in c slice.
	fmt.Println(c[2:]) // Include the index of 2 and slices up from index 2.
	// --------------
	// Initialize slice variable in single line.
	t := []string{"Ali", "Seylaneh", "Senior", "Software Engineer"}
	fmt.Println(t)
	// Looping over a collection like slice.
	for _, value := range t {
		fmt.Println(value)
	}
	for index, value := range t {
		fmt.Println(index, value)
	}
	// Using slice package utility example.
	t2 := []string{"Ali", "Seylaneh", "Senior", "Software Engineer"}
	if slices.Equal(t, t2) { // Equaling is compared by length of slice and elements
		fmt.Println("T2 and T are equal")
	}
	// Multi-dimensional slices
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j

		}
	}
	fmt.Println("2D: ", twoD)
}
