package main

func main() {
	//cards := newDeck()
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()

	// USING A FILE
	//cards.saveToFile("my_cards")
	//cards := newDeckFromFile(("my_cards"))
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
