package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}

	var i int = 5
	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
	}

	str := "Go is a beautiful language!"
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
}
