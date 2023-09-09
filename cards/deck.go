package main

import "fmt"

// A type which is a slice of strings
type deck []string

// Creates a new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Allows any deck to have access to this function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
