package main

import "fmt"

// A type which is a slice of strings
type deck []string

// Allows any deck to have access to this function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
