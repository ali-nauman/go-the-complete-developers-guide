package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
	// OR
	// contactInfo
}

func main() {
	// 3 ways to initialize structs
	//alex := person{"Alex", "Anderson"}
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//var alex person

	//alex.firstName = "Alex"

	//fmt.Println(alex)

	// Printing key/value pairs
	//fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim.party@foo.com",
			zipCode: 1234,
		},
	}

	jimPointer := &jim
	jimPointer.updateFirstName("John")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName(firstName string) {
	(*p).firstName = firstName
}
