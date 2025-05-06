package main

import "fmt"

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
	copy(c, v) // All the attributes of the slice is copied line cap and elements.
	fmt.Println(c, cap(c))
	v[2] = 33
	fmt.Println(v)
	fmt.Println(c)
	// ------------
	fmt.Println("Slicing in slice")
	fmt.Println(c[1:3], cap(c))
	fmt.Println(c[:4]) // Excluding the index of 4 in c slice.
}
