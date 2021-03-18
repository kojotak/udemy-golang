package main

import (
	"os"
	"testing"
)

//velke pismeno T na zacatku
func TestNewDeck(t *testing.T) { //t je test handler
	d := newDeck()

	//test poctu karet
	if len(d) != 13*4 {
		t.Errorf("deck has not enough cards, it has %v ", len(d)) //Errorf formatuje predane parametry
	}

	//test prvni karty
	if d[0] != "♠2" {
		t.Errorf("expected first card %v, but got %v", "♠2", d[0])
	}

	//test posledni karty
	if d[len(d)-1] != "♣A" {
		t.Errorf("expected last card %v, but got %v", "♣A", d[0])
	}
}

func TestSaveAndNewDeckFromFile(t *testing.T) { //dlouhe nazvy jsou "best practice" pro snazsi identifikaci
	fileName := "_deckTesting"
	os.Remove(fileName)
	deck := newDeck()
	deck.saveToFile(fileName)
	loaded := newDeckFromFile(fileName)

	if len(loaded) != 13*4 {
		t.Errorf("deck has not enough cards, it has %v ", len(loaded)) //Errorf formatuje predane parametry
	}

	os.Remove(fileName)

}
