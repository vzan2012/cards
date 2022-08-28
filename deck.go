package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Creating a new type of deck
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Func for Print the deck of cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal of Cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save file to hard disk
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// Read file from deck
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if(err != nil) {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s:= (strings.Split(string(bs), ","))

	return deck(s) 
}