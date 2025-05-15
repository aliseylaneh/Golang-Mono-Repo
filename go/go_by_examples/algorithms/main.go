package main

import "fmt"

func binarySearch(values []int, item int) int {
	low := 0
	high := len(values) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := values[mid]
		if item == guess {
			return mid
		} else if item < guess {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 10, 15, 18}
	fmt.Println(binarySearch(values, 6))
}
