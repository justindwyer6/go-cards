package main

import "fmt"

func main() {
	cards := newDeck()

	hand, cards := deal(cards, 5)
	fmt.Println("Hand:")
	hand.print()
	fmt.Println("Deck:")
	cards.print()
}
