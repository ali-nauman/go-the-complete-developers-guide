package main

import "fmt"

func main() {
	cards := []string{getNewCard(), getNewCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func getNewCard() string {
	return "Five of Diamonds"
}
