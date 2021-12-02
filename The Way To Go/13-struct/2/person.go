package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 1-struct as a value type
	var pers1 Person
	pers1.firstName = "Hello"
	pers1.lastName = "World"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	// 2-struct as pointer
	pers2 := new(Person)
	pers2.firstName = "Hello"
	pers2.lastName = "World"
	(*pers2).lastName = "World"
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	// 3-struct as literal:
	pers3 := &Person{"Hello", "World"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)
}
