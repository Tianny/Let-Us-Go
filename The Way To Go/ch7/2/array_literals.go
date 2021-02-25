package main

import "fmt"

func main() {
	var arrAge = [5]int{1,2,3,4,5}
	var arrLazy = [...]int{5,7,8,9}
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	var arrKeyValue2 = [...]string{3: "Chris", 4: "Ron"}

	for _, j := range(arrAge) {
		fmt.Println(j)
	}

	for _, j := range(arrLazy) {
		fmt.Println(j)
	}

	for _, j := range(arrKeyValue) {
		fmt.Println(j)
	}

	for _, j := range(arrKeyValue2) {
		fmt.Println(j)
	}
}