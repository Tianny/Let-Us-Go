package main

import "fmt"

type intVector []int

func (v  intVector) sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func main() {
	fmt.Println(intVector{1,2,3}.sum())
}