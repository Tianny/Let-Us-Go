package main

import "fmt"

func main() {
	var a int64 = 3
	var i interface{} = a

	v, ok := i.([]int)
	if !ok {
		fmt.Printf("%T\n", v)
		fmt.Printf("%d", v)
	}
}
