package main

import "fmt"

type twoInts struct {
	a int
	b int
}

func (tn *twoInts) AddThem()  int {
	return tn.a + tn.b
}

func (tn *twoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

func main() {
	two1 := new(twoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param  %d\n", two1.AddToParam(20))

	two2 := twoInts{3,4}
	fmt.Printf("The sum is %d\n", two2.AddThem())
}