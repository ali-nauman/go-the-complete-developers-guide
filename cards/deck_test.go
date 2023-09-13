package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	dLen := len(d)
	expectedLen := 16

	if dLen != expectedLen {
		t.Errorf("Expected deck length of %v, but got %v", expectedLen, dLen)
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs, but got %v", d[len(d)-1])
	}
}
