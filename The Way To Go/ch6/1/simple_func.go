package main

import "fmt"

func main() {
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", multiPly3Nums(2,5,6))
}

func multiPly3Nums(a int, b int, c int) int {
	return a * b * c
}