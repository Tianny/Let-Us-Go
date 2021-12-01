package main

import "fmt"

func modifySliceCorrect(innerSlice *[]string) {
	*innerSlice = append(*innerSlice, "a")
	(*innerSlice)[0] = "b"
	(*innerSlice)[1] = "b"
	fmt.Println(*innerSlice)
}

func main() {
	outerSlice := []string{"a", "a"}
	modifySliceCorrect(&outerSlice)
	fmt.Print(outerSlice)
}
