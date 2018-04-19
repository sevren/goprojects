package main

func main() {

	cards := newDeck()
	hand, _ := deal(cards, 3)
	hand.saveToFile("my_cards")

	cards = newDeckFromFile("my_cards")
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
