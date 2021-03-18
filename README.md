# udemy-golang

git repo for udemy's golang course

## section 2

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

homogenni pole, ktere se muze zvetsovat
deklaruje se jako nazev := [] type { obsah1, obsah2, ...}
pridavat se do slice muze pomoci append

### for cyklus

for index, promenna := range kolekce { ... }
pokud nechci pouzit index, tak musim misto nazvu dat _ (podtrzitko), jinak si bude compiler stezovat na deklarovanou a nepouzitou promennou

### type

GO neni objektove, daji se ale definovat vlastni typy

### receiver

neco jako staticka metoda, co patri k typu t a pridava mu funkcnost (funkci)
deklaruje se jako:

func (t typ) funkce() { ... }


