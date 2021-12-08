package main

func main() {
	//only in first nitialization : means that compiler needs to decide type based on the value assigned
	// card := "Ace of Spades"  //same as --> var card string = "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)
	//printState();

	// cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
