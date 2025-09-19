// The main package demonstrates how to construct, shuffle, and print a deck
// using the helper functions defined in deck.go. Treat this as a playground
// for experimenting with the API while learning Go.
package main

// main is the entry point of the program.
//
// Uncomment the sections below to explore other operations such as dealing
// cards, converting to a string, and saving to/reading from a file.
func main() {

	// cards := newDeck()

	// // Deal 5 cards from the top of the deck
	// // hand, remainingCards := deal(cards, 5)

	// // hand.print()
	// // remainingCards.print()

	// // Convert the deck to a single comma-separated string
	// // fmt.Println(cards.toString())

	// // Persist the deck to disk; check the file named "my_cards"
	// cards.saveToFile("my_cards")

	// // Read a deck previously saved to disk
	// cards := newDeckFromFile("my_cards")

	// cards.print()

	// Create, shuffle, and print a new deck
	cards := newDeck()
	cards.shuffle()

	cards.print()
}
