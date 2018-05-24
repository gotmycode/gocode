package main

import (
	"os"
	"testing"
)

//t is test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	//16 is the total number of 4 cardSuits x 4 cardValues
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but fot %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but go %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	//remove any remnants from past
	os.Remove("_decktesting")

	deck := newDeck()
	//should be called deck. (and not just d.)
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	//assert if correct length
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
