package main

import (
	"fmt"
)

func main() {
	callback(1, add)
}

func add(a, b int) {
	fmt.Printf("The sum of %d and %d is : %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}
