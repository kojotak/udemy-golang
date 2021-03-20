package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type czechBot struct{}

func main() {

	eb := englishBot{}
	cb := czechBot{}

	printGreeting(eb)
	printGreeting(cb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//lepsi by bylo mit pozdrav jako atribut ve strukture, tohle je jen pro demonstraci
func (englishBot) getGreeting() string {
	return "hello" //english specific
}

func (czechBot) getGreeting() string {
	return "nazdar" //czech specific
}
