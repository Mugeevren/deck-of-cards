package main

import "fmt"

//Create a new type of 'deck'
//which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//since we didn't use the index variables on for loops, we put underscore
	//so that go can userstand thet there is a variable there but I just don't care about it and not gonna use it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// usually reference with a one or two letter shortened version of the receiver argument
// unlike any other languages, to use a one or two letter to define a variable is ok for go
// instead of d we could use 'this' or 'self' and define the function like below,
//It would compile, but it is breaking convention, go avoids any method of 'this' or 'self'
// // func (this deck) print() {
// // 	for i, card := range this {
// // 		fmt.Println(i, card)
// // 	}
// // }
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
