package main

func main() {
	cards := newDeck()

	//carkou lze oddelit vice navratovych hodnot
	//hand, remaining := deal(cards, 5)
	//remaining.print()
	//hand.print()

	//vypise carkami oddelene seznam karet v baliku
	//fmt.Println(remaining.toString())

	//cards.saveToFile("mycards")
	//readCards := newDeckFromFile("mycards")
	//readCards.print()

	cards.shuffle()
	cards.print()
}
