package main

import "fmt"

type contactinfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactinfo
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
		contact: contactinfo{
			email:   "jim.party@foo.com",
			zipCode: 1234,
		},
	}
	fmt.Printf("%+v", jim)
}
