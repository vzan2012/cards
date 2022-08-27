package main

import "fmt"

func main() {

	// Slice Example
	cards := newDeck()

	// Print the deck of cards
	// cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// fmt.Println("----")
	// remainingCards.print()

	fmt.Println(cards.toString())

	cards.saveToFile("my-cards")

	// greeting := "Hi vzan2012"

	// fmt.Println([]byte(greeting))
}
