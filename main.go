package main

import "fmt"

func main() {
	cards := newDeck()

	// hand1 := deal()
	// fmt.Println("Hand:")
	// hand1.print()
	// fmt.Println("Deck:")
	// cards.print()

	hand, cards := deal(cards, 5)
	fmt.Println("Hand:")
	hand.print()
	fmt.Println("Deck:")
	cards.print()
}
