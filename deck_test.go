package main

import (
	// "fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 51 {
		t.Errorf("Error, actual len %v", len(d))
	}
}
