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
	var person1 Person
	person1.firstName = "Hello"
	person1.lastName = "World"
	upPerson(&person1)
	fmt.Printf("The name of the person is %s %s\n", person1.firstName, person1.lastName)

	// 2-struct as pointer
	person2 := new(Person)
	person2.firstName = "Hello"
	person2.lastName = "World"
	(*person2).lastName = "World"
	upPerson(person2)
	fmt.Printf("The name of the person is %s %s\n", person1.firstName, person1.lastName)

	// 3-struct as literal:
	person3 := &Person{"Hello", "World"}
	upPerson(person3)
	fmt.Printf("The name of the person is %s %s\n", person1.firstName, person1.lastName)
}
