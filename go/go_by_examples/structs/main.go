package main

import "fmt"

type Person struct {
	name        string
	phoneNumber int
}

func newPerson(name string, phoneNumber int) *Person {
	person := Person{name: "Ali", phoneNumber: 971933418093}
	return &person
}

func main() {
	person := newPerson("Ali Seylaneh", 9717989949)
	fmt.Println(&person)

	fmt.Println(person.name)
	fmt.Println(person.phoneNumber)
	// Accessing fields using point of the struct
	personTwo := &person
	fmt.Println(personTwo)
	// Single value anonymouse struct type
	dog := struct {
		name string
		kind string
	}{
		"Rex",
		"Golden Retriever",
	}
	fmt.Println(dog)
}
