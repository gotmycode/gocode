package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades"
	card := newCard()

	//card = "Five of Diamonds"
	fmt.Println(card)
}

func newCard() string {
	//when func is executed, func will return a value of type string (otherwise type empty)
	return "Five of Diamonds"
}
