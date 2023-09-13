package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// Converts a deck to a comma delimited string
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// Saves a deck to a file with the given name
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// Creates a new deck from a file with the given name
func newDeckFromFile(filename string) deck {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return strings.Split(string(bytes), ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
