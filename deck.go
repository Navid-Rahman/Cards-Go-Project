package main

import "fmt"

// Create a new type of 'deck'
// Which is a slice of Strings

type deck []string

func newDeck() deck {

	cards := deck{}

	cardTypes := []string{"Spades", "Diamond", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardType := range cardTypes {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardType)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
