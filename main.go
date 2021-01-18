package main

func main() {
	cards := newDeckFromFile("test-deck.csv")
	cards.shuffle()
	cards.print()
}
