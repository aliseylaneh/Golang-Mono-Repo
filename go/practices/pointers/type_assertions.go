package main

import "fmt"

type MyType struct {
	content string
}

func (t *MyType) describe() {
	if t == nil {
		fmt.Println("Value is nil")
		return
	}
	fmt.Println("My value is", t.content)
}

type MyInterface interface {
	describe()
}

func main() {
	//a := MyType{content: "Ali Seylaneh"}
	var a *MyType
	if b := a; b != nil {
		fmt.Println("B is not nil")
		return
	} else {
		fmt.Println("B is nil")
	}
	type MyString = string
}
