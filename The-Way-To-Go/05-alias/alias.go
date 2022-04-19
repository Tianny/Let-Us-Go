package main

import "fmt"

type myInt int

func main() {
	foo := myInt(1) // foo has type myInt
	i := int(foo)   // use type conversion to convert myInt to int
	fmt.Println(i)
}
