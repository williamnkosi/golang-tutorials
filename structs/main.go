package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson", contact: contactInfo{email: "email#gmaili.com", zipCode: 32098}}
	alexPointer := &alex
	fmt.Println(*alexPointer)
	alexPointer.updateName("Jim")
	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
