package main

import "fmt"

//Create a new type of 'deck'
//which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//****Underscore for Not Needed Variables*****
	//since we didn't use the index variables on for loops, we put underscore
	//so that go can understand that there is a variable there but I just don't care about it and not gonna use it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//***Unusual Naming Convention*****
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

//****Slice Range Syntax*****
//fruits := []string{"apple", "banana", "grape", "orange"}
//fruits[startIndexIncluding : upToNotIncluding] this format is build in go for easy use to fetch subset of an slice
//fruits[0:2] --> will bring "apple", "banana"
// one of the really interesting things about this kind of range syntax,
//it that we can opptionally leave out the numbers on either side of the column to have go automaticly infer that we want to start from the beggining or go all the way to the end of the slides
//ie: fruits[:2] is the same with fruits[0:2]
//ie: fruits[2:] ---> give me everything from the index of to to the very end of the slides --->

//****Multiple Return Values****
//go have support for returning multiple support from a function
//return two seperate values from this function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//check
// func (d deck) deal(handSize int) deck {

// }
