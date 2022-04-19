package main

import "fmt"

var num int = 10
var num2, num3 int

func main() {
	num2, num3 = getX2AndX3(num)
	printValues()
	num2, num3 = getX2AndX3t2(num)
	printValues()
}

func printValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, num2, num3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3t2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}
