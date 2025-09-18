package main

import "fmt"

func main() {

	// This is a Slice example
	cards := []string{newCard(), "Six of Diamond", "Two of Hearts"}

	cards = append(cards, "Five of Clubs")

	for i, card := range cards {

		fmt.Println(i, card)
	}

}

func newCard() string {

	return "Ace of Spades"
}
