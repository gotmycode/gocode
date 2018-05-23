package main

//for single package, no (), for multiple packages to import no separator needed
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of 'deck'
//which is a slice of strings
//in OO world, this new deck type borrows / extends behaviour of slice of type string
type deck []string

//annotate func to return a value of type 'deck'
//do you need a receiver? -- not (no need for a receiver)
func newDeck() deck {
	//for each suit in cardSuits, for each value in cardValues, add a new card of 'value of suit' to the 'cards' deck
	cards := deck{}
	//just slice of string (not deck)
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// replace variable i and j with _ (to fix declared and not used error message)
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// not put the name of func but a 'receiver'
//put on receiver of type deck and the function's name is 'Print'
// variable d is a reference to cards array -- d is cards, by convention 1st or 1st 3 letter of abbrev
// d is a value of type 'deck'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal function -- think about parameters
// 2 arguments passed to func -- can call this deal func with type deck and type int
func deal(d deck, handSize int) (deck, deck) {
	//go has support to show multiple values for one func -- in the above - both of type deck
	return d[:handSize], d[handSize:]
}

//should one set up a receiver or should the deck be passed in as an argument
//in example, we will set up a receiver
func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

//function will want to have receiver a type deck, so later it can be called by cards.saveToFile()
//does saveToFile require any arguments? func WriteFile will require a filename of type string, hence allow to pass arguments (instead of receiving)
//do we need a return value? func WriteFile can return error when attempting to write in harddrive
//perm 0666 anyone can read
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//func newDeckFromFile needs to be passed on argument filename to open
// the expectation is that one will receive a type deck hence we need to annotate deck
func newDeckFromFile(filename string) deck {
	//byteSlice of type error if nothing went wrong, it will have a value of nil
	bs, err := ioutil.ReadFile(filename)

	//check if err is not nil
	if err != nil {
		//Option #1 - log the error and return a call to newDeck()
		//Option #2 - log the error and entirely quit the program > for this example, we will take Option 2
		fmt.Println("Error:", err)
		//use func Exit from package os
		os.Exit(1)
	}

	//we have a byte slice and want to turn into a string
	//string(bs) // Ace of Spades, Two.....
	//return a slice of strings
	s := strings.Split(string(bs), ",")

	//conversion from slice of string to deck type
	return deck(s)
}

//func signature - we want to call cards.shuffle at some point
//we will need a receiver (take the deck and randomize) nothing to return
func (d deck) shuffle() {
	//seed value = source
	//pass some value to NewSource > use time package
	//func Now & func (Time)UnixNano >every single time we run program we will use a new time
	source := rand.NewSource(time.Now().UnixNano())

	//type Rand
	r := rand.New(source)

	//for index or simply i (per convention) -- we need to get the element that we need to iterate over
	for i := range d {
		//use rand function in math package -- pseudo random generator (by default go uses the exact same seed value)
		//len is the length of slice
		//newPosition := rand.Intn(len(d) - 1)

		//with fix to the same seed value (src) -> func New(src Source) *Rand
		newPosition := r.Intn(len(d) - 1)

		//swap elements, take whatever is at newPosition and assign to i, take whatever is in i and assign to newPosition
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
