package main

import "fmt"

func main() {
	var a =  [...]string{"a", "b", "c", "d"}
	for i, j := range a {
		fmt.Println("Array item", i, "is", j)
	}
}