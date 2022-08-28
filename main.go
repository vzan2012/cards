package main

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

	cards.print()

	// fmt.Println([]byte(greeting))
	// fmt.Println(newDeckFromFile("my-cardss"))
}
