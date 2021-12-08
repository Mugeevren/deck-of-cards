package main

func main() {
	//only in first nitialization : means that compiler needs to decide type based on the value assigned
	// card := "Ace of Spades"  //same as --> var card string = "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)
	//printState();

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	deck.print(cards)
}

func newCard() string {
	return "Five of Diamond"
}