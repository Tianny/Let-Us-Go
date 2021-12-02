package main

import (
	"fmt"
)

func main() {
	// make an add2 function, give it a name p2, and call it:
	p2 := add2()
	fmt.Printf("Call Adder2 for 3 gives: %v\n", p2(3))
	// make a special adder function, a gets value 2:
	TwoAdder := adder(2)
	fmt.Printf("The result is %v\n", TwoAdder(3))
}

func add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
