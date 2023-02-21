package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3, 4, 5}
	fmt.Println("original a=", a)

	for i := range a[:] {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
	}
	fmt.Println("after for range loop, a = ", a)
}
