package main

func main() {
	cards := deck{getNewCard(), getNewCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func getNewCard() string {
	return "Five of Diamonds"
}
