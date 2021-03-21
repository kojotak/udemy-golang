# udemy-golang

git repo for udemy's golang course

## poznamky z kurzu

### build + spusteni

go run main.go

### pouze build bez spusteni

go build main.go

### package main

obsahuje veci ke spusteni, ne knihovny a podpurny kod. Musi byt uvedeno na zacatku souboru nehlede na adresarovou strukturu

### deklarace promenne

zkracena forma - operator :=, jinak se musi zadat: var nazev typ

### deklarace navratoveho typu funkce

se uvadi mezi () a {}, nikoliv na zacatku

### slice

* homogenni pole, ktere se muze zvetsovat
* deklaruje se jako nazev := [] type { obsah1, obsah2, ...}
* pridavat se do slice muze pomoci append

### for cyklus

* for index, promenna := range kolekce { ... }
* pokud nechci pouzit index, tak musim misto nazvu dat _ (podtrzitko), jinak si bude compiler stezovat na deklarovanou a nepouzitou promennou

### type

GO neni objektove, daji se ale definovat vlastni typy

### receiver

neco jako staticka metoda, co patri k typu t a pridava mu funkcnost (funkci)

deklaruje se jako:

func (t typ) funkce() { ... }

### slice range

[0:2]

* od 0 vcetne (coz je prvni prvek), do 2 exkluzivne, tj. vrati 2 kusy
* x: znamena vrat od x do konce
* :x znamena vrat od nuly do x

### zmena typu

* castovani pomoci zavorek, napriklad string na byte slice:  []byte("retezec")
* (ulozeni dat na disk - nejdriv prevest na string, pak na byte slice)

### testovani

* nazev testovaciho souboru musi koncit _test.go
* testy se spousti prikazem: go test
* nezobrazuje se pocet uspesnych testu, jen chyby

### struktury

* = objekty bez metod
* pri deklaraci jen popis casti - kazdy na novy radek
* pri inicializaci jen slozene zavorky a carkou oddeleny seznam casti
* to muze byt ale matouci, nebot zavisime na poradi. Lepsi je seznam dvojic - nazev:hodnota
* pri deklaraci bez inicializace (var nazev typ) se priradi automaticky "zero value" (namisto null)
* pokud inicializuji struct pomoci {}, musim dat oddelovac , i za posledni polozku
* pokud deklaruji struct uvnitr struct, nemusim uvadet nazev promenne, vytvori se automaticky podle typu structory
* pokud ve funkci zmenim hodnotu struct, tak se to zvenku neprojevi, a je potreba pouzit...

### ukazatele

* go je "pass by value" - pri volani funkce se zkopiruji a predaji zkopirovana data, nikoliv ukazatel
* operator &promenna - vrati adresu, kam nazev promenne ukazuje
* operator *ukazatel - vrati hodnotu, na kterou ukazatel ukazuje
* pozor, pri deklaraci funkce s receiverem zapis (pointer *typ) rika, ze promenna pointer je typu "ukazatel na typ"
* ale v tele funkce s receiverem zapis *pointer znamena "hodnota, kam ukazuje promenna pointer"
* zkratka - go si umi ukazatel domyslet, takze neni potreba vzdycky nejprve vyrabet ukazatel pomoc &
* co kdyz misto struct predam slice? NEpouzije se "pass by value"
* value types: primitivni typy + stringy + struct
* reference type: slice + kolekce + channels + pointers + functions
* (u referencenich typu nemusime ukazatele vubec resit, tj. jsou potreba jen u struktur)

### mapy

* klice i hodnoty jsou staticky typovane
* syntax: map[typKlice]typHodnoty
* alternativni vytvoreni prazdne mapy: make(map[int]string)
* mazani: delete(mapa, klic)
* oproti strukture je referencni typ, musi mit stejne typy klicu a hodnot, muze se iterovat pres klice, ktere jsou dynamicke (mohou se pridavat/odebirat)

### interfaces

* trosku jiny vyznam nez v OOP: pokud dana struktura obsahuje vsechny deklarovane funkce, potom rikame, ze je i daneho typu (rozhrani). Rikame tedy, ze rozhrani je "implicitni" (oproti napr. jave, kde se musi explicitne uvest implements)
* tj. vlastni typy nerikaji, ze maji nejaky vztah k rozhrani - to vyplyva z toho, jestli obsahuji dane funkce
* interface je tedy definovan mnozinou funkci
* i kdyz je interface type, nelze z nej ale vytvorit value (instanci)
* nemohou byt genercike. Go vlastne generiky vubec nema

### routines 

* routine - v cem bezi program, neco jako vlakno
* kazdy program bezi minimalne v jedne routine s nazvem main
* kdyz chci pustit neco v nove rutine, pridam pred volanim funkce: go
* po zavolani go pokracuje puvodni routine dal
* takove routine se rika "child"
* go scheduler planuje go routiny i na jednom jadru procesoru, ale defaultne 1 routina na 1 jadro
* concurrency - bezi nekolik routine najednou na jednom jadru procesoru
* paralellismus - bezi nekolik routine najednou na vice jadrech procesoru
* ! problem - co kdyz main routine dobehne driv, nez children routines?

### channels

* channel - umoznuje komunikovat mezi routinama. Jde o jediny zpusob komunikace mezi routinama
* channel ma typ, ktery omezuje, co si routiny mezi sebou vymenuji
* posilani a prijimani zprav pres channel - operator <- (sipka)
* tj. poslani: channel <- hodnota //posle hodnotu do channelu
* a prijeti: promenna <- channel //ceka na zpravu z channelu - blokujici operace!
