package main

import "fmt"

func main() {
	cards := newDeck()

	hand1 := cards.deal()
	fmt.Println("Hand:")
	hand1.print()
	fmt.Println("Deck:")
	cards.print()
}
