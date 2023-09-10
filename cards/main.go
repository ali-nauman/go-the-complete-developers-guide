package main

import "fmt"

func main() {
	cards := newDeck()
	err := cards.saveToFile("cards.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	cards2 := newDeckFromFile("cards2.txt")
	cards2.print()
}
