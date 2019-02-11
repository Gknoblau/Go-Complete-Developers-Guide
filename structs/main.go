package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Jeff",
		contactInfo: contactInfo{
			email:   "gim@gmail.com",
			zipCode: 12345,
		},
	}

	jim.updateName("Jimmy")
	jim.print()

}

func (p *person) updateName(newFirstname string) {
	(*p).firstName = newFirstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
