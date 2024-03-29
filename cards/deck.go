package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck'
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

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	 return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0644)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)

	if err != nil {
		// Option 1: Log error, and create and return a newDeck()
		// Option 2: Log error, and quit the program 
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(byteSlice), ","))
}

func (d deck) shuffle() {
	rand.Shuffle(len(d), func(i int, j int) {
		d[i], d[j] = d[j], d[i]
	})
}