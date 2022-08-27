package main

import "fmt"

func main() {

	// Slice Example
	cards := newDeck()

	// Print the deck of cards
	// cards.print()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	fmt.Println("----")
	remainingCards.print()
}
