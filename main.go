package main

func main() {
	cards := newDeck()
	cards.saveToFile("test-deck.csv")
}
