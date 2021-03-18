# udemy-golang

git repo for udemy's golang course

## section 3 

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