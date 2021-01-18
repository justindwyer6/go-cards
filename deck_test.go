package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	a := 16

	if len(d) != a {
		t.Errorf("Expected deck length of %v, but got %v", a, len(d))
	}
}
