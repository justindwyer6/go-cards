package main

func main() {
	cards := newDeckFromFile("test-deck.csv")
	cards.print()
}
