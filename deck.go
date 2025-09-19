// Package main contains a simple playing card deck implementation used
// for learning core Go concepts such as slices, methods, receivers,
// multiple return values, file I/O, and testing.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// deck represents a collection of playing cards where each card is a string
// like "Ace of Spades". It is implemented as a slice of strings so we can
// easily append, slice, iterate, and pass by reference semantics.
type deck []string

// newDeck creates a new deck with a fixed set of suits and values.
//
// It returns a deck containing 16 cards (4 values x 4 suits) in a stable
// order. This is helpful for testing and predictable behavior.
func newDeck() deck {

	cards := deck{}

	cardTypes := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardType := range cardTypes {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardType)
		}
	}

	return cards
}

// print writes each card in the deck to standard output, prefixed with its
// index position. Useful for quick debugging while learning.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal splits the provided deck into two parts:
// - the first return value is the "hand" containing handSize cards
// - the second return value is the remainder of the deck
//
// Note: This function does not mutate the original deck slice passed in;
// it returns new slices that reference the same underlying array.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// toString joins all cards in the deck into a single comma-separated string.
// This string format is used when saving to and reading from a file.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saveToFile writes the deck to the given file using 0666 permissions.
//
// It serializes the deck with toString() and stores the bytes on disk.
// The returned error should be checked by callers when used in real apps.
func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

// newDeckFromFile reads a deck from a file previously saved by saveToFile.
//
// If an error occurs (such as a missing file), it prints the error and
// exits the program with status code 1. For production code, prefer
// returning the error instead of exiting.
func newDeckFromFile(fileName string) deck {

	bs, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")

	return deck(s)

}

// shuffle randomizes the order of cards in the deck in place.
//
// Note: This uses math/rand without an explicit seed, which means the
// sequence will be deterministic unless a seed is set elsewhere
// (e.g., rand.Seed(time.Now().UnixNano())). Additionally, using
// rand.Intn(len(d) - 1) avoids selecting the last index and introduces
// slight bias; a better approach is rand.Intn(len(d)).
func (d deck) shuffle() {

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
