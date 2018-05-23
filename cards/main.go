package main

func main() {
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")

	//cards := newDeck()
	// for index i with the current card we're iterating take the slide of cards and loop over it
	// in for loops we re throwing away any previous declaration, we will re-initialize them
	//for i, card := range cards {
	//run this one time for each card in the slice
	//	fmt.Println(i, card)
	//}

	//calling the func in deck.go
	//cards.print()

	//hand, remainingDeck - 1st return value assigned to hand, 2nd return value to remainingDeck
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()

	//type conversion example
	//greeting := "Hi there!"
	//fmt.Println([]byte(greeting))

	//cards := newDeck()
	//to print on screen the joined strings
	//fmt.Println(cards.toString())

	//attempt to save to hard drive
	//cards.saveToFile("my_cards")

	//attempt to read a file that does not exist (my), correct filename is my_cards
	//cards := newDeckFromFile("my")
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

//func newCard() string {
//when func is executed, func will return a value of type string (otherwise type empty)
//	return "Five of Diamonds"
//}
