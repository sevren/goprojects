package main

import (
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 , but got %v", len(d))
	}

	if strings.Compare(d[0], "Ace of Spades") != 0 {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}

	if strings.Compare(d[len(d)-1], "Four of Clubs") != 0 {
		t.Errorf("Expected last card of Four of Clubs but got %v", d[len(d)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck go %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
