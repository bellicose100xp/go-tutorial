package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d)) 
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of 'Ace of Spades', but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of 'Four of Clubs', but got %v", d[len(d) - 1])
	}
}

func TestSaveDeckToFileAndCreateNewDeckFromSavedFile(t *testing.T) {
	filename := "_cards"

	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in the deck, but got %v", len(loadedDeck))
	}

	os.Remove(filename)
}