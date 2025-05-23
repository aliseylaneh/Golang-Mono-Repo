package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// Arrays
	var x = [...]int{23, 123, 12, 44, 5}
	x[4] = 54
	// Multi dimensional arrays
	var y = [2][3]int{}
	y[0] = [3]int{23, 22, 41}
	y[1] = [3]int{23, 22, 41}
	// Slices
	sliceX := []int{23, 11, 3321, 11}
	sliceY := []int{223, 122, 331, 22}
	fmt.Println(slices.Equal(sliceX, sliceY))
	sliceXY := []int{1, 2, 3, 4}
	sliceYY := []int{1, 2, 3, 4}
	fmt.Println(slices.Equal(sliceXY, sliceYY))
	// Maps
	var nilMap map[string]int
	//nilMap["1"] = 1 // writing in nil map cause panic
	fmt.Println(nilMap)
	// The comma ok idiom
	customers := map[string]string{"A": "Ali", "B": "Jake"}
	providers := map[string]string{"A": "Ali", "B": "Jake"}
	fmt.Println(maps.Equal(customers, providers))
	value, exists := customers["A"]
	fmt.Println(value, exists)
	value, exists = customers["C"]
	fmt.Println(value, exists)
	delete(customers, "B")
	fmt.Println(customers)
	clear(customers)
	// Anonymous struct
	pet := struct {
		name string
		kind string
	}{name: "Flip", kind: "Cats"}
	fmt.Println(pet)
}
