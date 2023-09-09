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

// Prints each card in a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deals a hand, returning it and the remaining deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
