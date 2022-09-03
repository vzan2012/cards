package main

import "fmt"

func main() {

	// Slice Example
	// cards := newDeck()

	// Print the deck of cards
	// cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// fmt.Println("----")
	// remainingCards.print()

	// fmt.Println(cards.toString())

	// cards.saveToFile("my-cards")

	// greeting := "Hi vzan2012"

	cards := newDeckFromFile("my-cards")

	// Print the cards
	cards.print()

	// Shuffle the cards
	cards.shuffle()

	// Print the cards
	fmt.Println("**********************")

	cards.print()



	// fmt.Println([]byte(greeting))
	// fmt.Println(newDeckFromFile("my-cardss"))
}
