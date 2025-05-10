package main

import "fmt"

type functionAsArg func(string) string

func wordPrinter(word string) string {
	return fmt.Sprintf("The accepted word is -> %v", word)
}
func myFunction(word string, function functionAsArg) {
	returnedValue := function(word)
	fmt.Println(returnedValue)
}

// Variadic function
func variadicFunction(names ...any) {
	for i, name := range names {
		switch valueType := name.(type) {
		case string:
			fmt.Printf("The index %v is a %T and the value is %v\n", i, valueType, name)
		case int:
			fmt.Printf("The index %v is an %T and the value is %v\n", i, valueType, name)
		case float32:
			fmt.Printf("The index %v is a %T and the value is %v\n", i, valueType, name)
		default:
			fmt.Printf("The type of index %v is a %T and the value is %v, and it's not listed in our system\n", i, valueType, name)
		}
	}
}

// Multiple return type for function
// This function returns the hit target of a coin for the relevant day earned sale.
// Also, it will bring back if it was eligible to be a progress sale for that coin.
func eligibleTarget(targetSales []int, dayEarnedSale int) (int, bool) {
	if targetSales[0] > dayEarnedSale {
		return 0, false
	}
	var winner int
	for i, targetSale := range targetSales {
		if i == len(targetSales)-1 {
			winner = targetSale
		}
		if targetSale <= dayEarnedSale && dayEarnedSale < targetSales[i+1] {
			winner = targetSale
			return winner, true
		}
	}
	return winner, true
}
func main() {
	// Multiple return value function
	targetSales := []int{1200, 1400, 1550, 1600}
	fmt.Println(eligibleTarget(targetSales, 1500))
	// Function as args
	myFunction("Ali Seylaneh", wordPrinter)

	// Variadic functions
	variadicFunction(3, 4, "Ali Seylaneh", float32(3.5), 4.6)
}
