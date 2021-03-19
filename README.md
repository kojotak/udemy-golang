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