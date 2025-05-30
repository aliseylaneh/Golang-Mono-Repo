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

type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst *List[T]) AllElements() []T {
	var elements []T
	for e := lst.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}
	return elements
}
func main() {
	s := []int{1, 2, 3, 5, 6}
	v := 5
	fmt.Println(SlicesIndex[[]int, int](s, v))
	fmt.Println(SliceIndexBinarySearch(s, v))

	linkedList := List[string]{}
	linkedList.Push("Ali")
	linkedList.Push("Seylaneh")
	linkedList.Push("Mahan")
	fmt.Println(linkedList.AllElements())
}
