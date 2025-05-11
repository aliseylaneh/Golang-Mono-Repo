package main

import "fmt"

func printerOfReference(value *string) {
	fmt.Println(*value)
	*value = "Seylaneh"
	fmt.Println(value)

}

func main() {
	var value string = "Ali Seylaneh"
	printerOfReference(&value)
	fmt.Println(value)
	fmt.Println(&value)
}
