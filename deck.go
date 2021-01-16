package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// func (d deck) deal() deck {
// 	var newHand []string
// 	for i, card := range d {
// 		if i < 5 {
// 			// append 1 card to newHand
// 			newHand = append(newHand, card)
// 			// remove that card from deck
// 			// Taken from https://yourbasic.org/golang/delete-element-slice/
// 			d[i] = d[len(d)-1]
// 			d[len(d)-1] = ""
// 			d = d[:len(d)-1]
// 		}
// 	}
// 	return newHand
// }

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
