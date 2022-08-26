package main

import "fmt"

func main() {

	// Ways of declaring variable
	cardNew := newCard()

	// Slice Example
	cards := []string{"Ace of Diamonds", newCard()}

	cards = append(cards, "Six of Spades")

	fmt.Println(cardNew)

	fmt.Println(cards)

	// Loop Iteration
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

// Initialize the variables
func newCard() string {
	return "Five of Diamonds"
}
