package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if num := rand.Intn(10); num%2 > 0 {
		fmt.Println("The 'BaghiMandeh' is bigger than 0")
	} else {
		fmt.Println("The condition is not standing")
	}
}
