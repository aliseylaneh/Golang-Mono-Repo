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
func main() {
	myFunction("Ali Seylaneh", wordPrinter)
}
