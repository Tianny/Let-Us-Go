// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly.
package main

import (
	"LetUsGo/Go-In-Action/chapter5/listing74/entities"
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Create a value of type Admin from the entities package.
	a := entities.Admin{
		Rights: 10,
	}

	// Set the exported fields from the unexported
	// inner type.
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}
