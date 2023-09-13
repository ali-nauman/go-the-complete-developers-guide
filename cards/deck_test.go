package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	dLen := len(d)
	expectedLen := 16

	if dLen != expectedLen {
		t.Errorf("Expected deck length of %v, but got %v", expectedLen, dLen)
	}
}
