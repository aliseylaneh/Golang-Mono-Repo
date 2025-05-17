package main

import "fmt"

type User struct {
	name string
	role string
	ID   int
}

func createUser(name string, role string, nationalId int) *User {
	user := User{name: name, role: role, ID: nationalId}
	return &user
}

func (user *User) printOutUser() {
	fmt.Println(user.name, user.role, user.ID)
}
func (user User) printOutUserNoPointer() {
	fmt.Println(user.name, user.role, user.ID)
}
func main() {
	user := createUser("Ali", "Software Engineer", 23342)
	user.printOutUser()
	user.printOutUserNoPointer()
}
