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
	targetSales := []int{1200, 1400, 1550, 1600}
	fmt.Println(eligibleTarget(targetSales, 1500))
	//myFunction("Ali Seylaneh", wordPrinter)
}
