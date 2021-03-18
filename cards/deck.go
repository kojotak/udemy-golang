package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	suits := []string{"♠", "♦", "♥", "♣"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, suit+""+value)
		}
	}

	return cards
}

//rozdeli balicek na dve casti - opet typu deck
//neni tu receiver, ale deck je parametr, aby bylo jasne, ze funkce deal nemodifikuje balicek,
//   ale vraci ruku a zbytek
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	slices := []string(d)            //deck je stejne definovany jako typ []string
	return strings.Join(slices, ",") //spojim carkou
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteslice, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("fatal error: ", error)
		os.Exit(42)
	}
	commaSeparated := string(byteslice) //pretypovani z bytu
	stringslice := strings.Split(commaSeparated, ",")
	return deck(stringslice) //pretypovani, protoze  deck je slice of string
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) //zdroj nahodnosti na zakladne aktualniho casu
	random := rand.New(source)                      //poskytuje funkce z daneho zdroje nahodnosti

	for i := range d { //staci mi index, nepotrebuju hodnotu
		//newPosition := rand.Intn(len(d) - 1) //vicenasobne spusteni povede ke stejnemu vysledku
		//potrebujeme nahodny seed, pouzijeme typ Rand - viz zacatek teto funkce
		newPosition := random.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i] //jednoradkovy swap
	}
}
