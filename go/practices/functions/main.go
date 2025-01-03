package main

import "fmt"

type useCaseFunctionArgs func(int, string) (int, string, string)

func handleUploadUseCase(id int, name string, useCaseFunction useCaseFunctionArgs) {
	id, name, typeOfStruct := useCaseFunction(id, name)
	fmt.Printf("id: %v, name: %v, Type of Struct: %v\n", id, name, typeOfStruct)
}

type Product struct {
	id    int
	name  string
	color string
}

type Journal struct {
	id     int
	name   string
	number int
}

func journalUseCase(id int, name string) (int, string, string) {
	journal := Journal{id: id, name: name, number: 3321}
	return journal.id, journal.name, "Journal"
}

func productUseCase(id int, name string) (int, string, string) {
	product := Product{id: id, name: name, color: "red"}
	return product.id, product.name, "Product"
}
func useCase(structName string) (useCaseFunctionArgs, error) {
	switch structName {
	case "product":
		return productUseCase, nil
	case "journal":
		return journalUseCase, nil
	default:
		return nil, fmt.Errorf("faild to find the usecase")
	}

}
func myFunction(name string, phoneNumbers ...int) {
	fmt.Println(name, phoneNumbers)
}

func namedReturnedValues(name, address string) (processedName, processedAddress string) {
	processedName, processedAddress = name, address
	return // Remember to never use an empty return while you have named return value, because when you do this it will
	// return the latest assigned value to those named return values.
}
func f1(name string) string {
	return name
}
func valueFunctions() {
	// You can declare functions as value and then assign a function to that variable which is defined as function.
	var useCaseFunctionArgsTwo func(string) string // zero value for a function is nil. Attempting to run a function
	// with nil value will cause panic
	useCaseFunctionArgsTwo = f1
	fmt.Println(useCaseFunctionArgsTwo("Hello"))
}
func anonymousFunction() {
	f := func(name string) {
		fmt.Println(name)
	}
	f("Ali")
	defer f("A")           // Anonymous functions are useful when it comes to defer and goroutines.
	defer valueFunctions() // You can defer multiple time for multiple functions, but it's gonna be LAST IN FIRST OUT.
}
func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first:", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("second:", val)
	}(a)
	a = 30
	fmt.Println("exiting:", a)
	return a
}
func main() {
	myFunction("Ali", 12312, 123123, 123123)
	phoneNumbers := []int{123, 123, 123}
	myFunction("Ali", phoneNumbers...)
	journalFunctionUseCase, err := useCase("journal")
	if err != nil {
		return
	}
	productFunctionUseCase, err := useCase("product")
	if err != nil {
		return
	}
	handleUploadUseCase(10, "Journal of New York", journalFunctionUseCase)
	handleUploadUseCase(24, "Neapolitan Pizza", productFunctionUseCase)
	name, address := namedReturnedValues("Ali", "New York")
	fmt.Println(name, address)
	valueFunctions()
	anonymousFunction()
	a := deferExample()
	fmt.Println(a)
}
