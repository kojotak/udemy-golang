package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Bean",
		contactInfo: contactInfo{
			email:   "jim@bean.cz",
			zipcode: "123456",
		},
	}
	//pass by value - nezafunguje
	jim.updateNaiveName("jim-u1")
	jim.print() // ! vypise jim, nikoliv joe podle ocekavani

	//oprava pomoci ukazatelu
	jimPointer := &jim
	jimPointer.updateName("jim-u2")
	jimPointer.print()

	//zkraceny zapis pomoci ukazatelu
	jim.updateName("jim-u3") //nesedi typu s funkci updateName! GO si ukazatel automaticky vyrobi
	jim.print()
}

func (pointer *person) updateName(newFirstName string) {
	(*pointer).firstName = newFirstName
}

func (p person) updateNaiveName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)

}
