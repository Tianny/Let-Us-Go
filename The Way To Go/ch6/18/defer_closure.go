package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer func ()  {
			fmt.Print(i)
		}()
	}

	for i := 0; i < 3; i++ {
		defer fmt.Print(i)
	}

}