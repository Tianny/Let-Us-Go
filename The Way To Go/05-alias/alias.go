package main

import "fmt"

type myint int

func main() {
	foo := myint(1) // foo has type myint
	i := int(foo)   // use type conversion to convert myint to int
	fmt.Println(i)
}
