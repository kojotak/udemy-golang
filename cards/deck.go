package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	suits := []string{"spades", "diamonds", "hears", "clubs"}
	values := []string{"ace", "two", "three"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, suit+" "+value)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
