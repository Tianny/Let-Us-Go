package main

import "fmt"

func main() {
	a := [10]int{}
	a[0] = 2
	fmt.Println(a)

	p := new([10]int)
	p[0] = 1
	fmt.Println(p)
}