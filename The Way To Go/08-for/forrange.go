package main

import "fmt"

func main() {
	str := "Go is a beautiful language!"
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
}
