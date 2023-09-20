package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// 3 ways to initialize structs
	//alex := person{"Alex", "Anderson"}
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	var alex person

	alex.firstName = "Alex"

	fmt.Println(alex)

	// Printing key/value pairs
	fmt.Printf("%+v", alex)
}
