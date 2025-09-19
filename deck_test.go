// Tests for the deck implementation. Run with `go test`.
package main

import (
	"os"
	"testing"
)

// TestNewDeck verifies that newDeck constructs a deck with the expected
// number of cards and preserves a stable order for the first and last cards.
func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

// TestSaveToDeckAndNewDeckFromFile ensures a deck can be serialized to disk
// and reconstructed from that file. It uses a throwaway test filename.
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove(("_deckTesting"))

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
