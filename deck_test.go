package main

import (
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