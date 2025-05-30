package main

import "fmt"
import constraints "golang.org/x/exp/constraints"

// SlicesIndex This is linear search and include generic type, with time complexity of O(n)
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}
	return -1
}

// SliceIndexBinarySearch This is binary search implemented by generic type using constraints library.
// With time complexity of O(log n)
func SliceIndexBinarySearch[S ~[]E, E constraints.Ordered](s S, v E) int {
	leftIndex := 0
	rightIndex := len(s) - 1
	for i := leftIndex; i < rightIndex; i++ {
		middleIndex := (leftIndex + rightIndex) / 2
		if s[middleIndex] == v {
			return middleIndex
		} else if v < s[middleIndex] {
			rightIndex = middleIndex - 1
		} else {
			leftIndex = middleIndex + 1
		}
	}
	return -1
}

func main() {
	s := []int{1, 2, 3, 5, 6}
	v := 5
	fmt.Println(SlicesIndex(s, v))
	fmt.Println(SliceIndexBinarySearch(s, v))
}
