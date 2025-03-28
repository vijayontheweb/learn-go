package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}

func TestWriteToFileAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.writeToFile("_decktesting")
	readDeck := stringToDeck(readFromFile("_decktesting"))
	if len(readDeck) != 52 {
		t.Errorf("Expected 52 but found %v", len(readDeck))
	}
	os.Remove("_decktesting")
}
