package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r rect) perim() float64 {
	return (r.height + r.width) * 2
}
func (r rect) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func printPerim(shape geometry) {
	if c, ok := shape.(circle); ok {
		fmt.Printf("shape is %T, perm is: %v\n", c, shape.perim())
	} else if c, ok := shape.(rect); ok {
		fmt.Printf("shape is %T, perm is: %v\n", c, shape.perim())
	} else {
		fmt.Println("No relevant type of geometry is passed for area calculation.")
	}
}
func printArea(shape geometry) {
	if c, ok := shape.(circle); ok {
		fmt.Printf("shape is %T, area is: %v\n", c, shape.perim())
	} else if c, ok := shape.(rect); ok {
		fmt.Printf("shape is %T, area is: %v\n", c, shape.perim())
	} else {
		fmt.Println("No relevant type of geometry is passed for area calculation.")
	}

}

func main() {
	rectangle := rect{width: 30, height: 10}
	printArea(rectangle)
	printPerim(rectangle)
	circular := circle{radius: 30}
	printArea(circular)
	printPerim(circular)
}
