package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://golang.org",
		"http://stackoverflow.com",
	}

	//checkLinksSequentially(links)
	//checkLinksConcurrently(links)
	checkLinksRepeatedly(links)
}

//////////////////////////////////////////////

func checkLinksSequentially(links []string) {
	fmt.Println("sequential check...")
	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "may be down!")
	} else {
		fmt.Println(link, "is OK")
	}
}

//////////////////////////////////////////////

func checkLinksConcurrently(links []string) {
	fmt.Println("concurrent check...")
	//vytvori channel typu string
	c := make(chan string)

	for _, link := range links {
		//bez go pobezi vsechny kontroly sekvencne
		go checkLinkWithChannel(link, c)
	}

	//fmt.Println(<-c) //vypise pouze prvni zpravu z channelu, neceka na ostatni

	for i := 0; i < len(links); i++ {
		//pocka na vsechny zpravy, ktere maji child routine poslat
		fmt.Println(<-c)
	}
}

func checkLinkWithChannel(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link + " may be down!"
	} else {
		c <- link + " is OK"
	}
}

///////////////////////////////////////////////////////

func checkLinksRepeatedly(links []string) {
	fmt.Println("repeatedly concurrent check...")
	c := make(chan string)

	for _, link := range links {
		go checkLinkWithChannelRepeatedly(link, c)
	}

	for l := range c { //nekonecna smycka, slo by zapsat jenom jako for {
		//pusti znova kontrola toho linku, ktery prave dobehnul
		//time.Sleep(time.Second) // neni dobry napad davat sleep do main routine
		go func(link string) {
			time.Sleep(time.Second) //nechceme davat sleep do checkLinkWithChannelRepeatedly
			//protoze by to zbrzdilo i prvni volani
			checkLinkWithChannelRepeatedly(link, c)
		}(l) //zavorky a predani linku musi byt na konci, aby se funkce hned spustila
		//kdybychom uvnitr child routine pouzili l, tak dostanene varovani kompilatoru,
		// protoze bychom pristupovali na promennou z jine routiny
	}
}

func checkLinkWithChannelRepeatedly(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "may be down!")
	} else {
		fmt.Println(link, "is OK")
	}
	c <- link //posle zpravu, ze jsme s timto odkazem hotovy
}
