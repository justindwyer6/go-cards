package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	lengthAssertion := 16
	if len(d) != lengthAssertion {
		t.Errorf("Expected deck length of %v, but got %v", lengthAssertion, len(d))
	}

	firstCardAssertion := "Ace of Spades"
	if d[0] != firstCardAssertion {
		t.Errorf("Expected first card to be %v, but got %v", firstCardAssertion, d[0])
	}

	lastCardAssertion := "Four of Clubs"
	if d[len(d)-1] != lastCardAssertion {
		t.Errorf("Expected last card to be %v, but got %v", lastCardAssertion, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	lengthAssertion := 16
	if len(loadedDeck) != lengthAssertion {
		t.Errorf("Expected deck length of %v, got %v", lengthAssertion, len(loadedDeck))
	}

	os.Remove("_decktesting")
}
