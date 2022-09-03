package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	expectedCards := 16
	expectedFirstCard := "Ace of Spades"
	expectedLastCard := "Four of Clubs"

	d := newDeck()
	
	if(len(d)!=expectedCards) {
		t.Errorf("Expected deck length of 16 but received %v", len(d))
	}

	if(d[0] != expectedFirstCard) {
		t.Errorf("Expected first card to be %v, but got %v", expectedFirstCard, d[0])
	}

	if(d[len(d)-1] != expectedLastCard) {
		t.Errorf("Expected last card to be %v, but got %v", expectedLastCard, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	expectedCards := 16
	testFileName := "_decktesting"

	// Remove any test file exists 
	os.Remove(testFileName)

	// Call the newDeck func 
	d := newDeck()

	// Save the deck to testfile
	d.saveToFile(testFileName)

	// Load the Deck testfile
	loadedDeck := newDeckFromFile(testFileName)

	// Test case
	if(len(loadedDeck) != expectedCards) {
		t.Errorf("Expected %v cards in deck, but got %v", expectedCards, len(loadedDeck))
	}

	// Remove the file 
	os.Remove(testFileName)
}