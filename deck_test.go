package main

import "testing"

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
