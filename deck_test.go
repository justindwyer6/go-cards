package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	e := 15

	if len(d) != e {
		t.Errorf("Expected deck length of %v, but got %v", e, len(d))
	}
}
