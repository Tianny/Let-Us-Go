package main

import "fmt"

func modifySlice(innerSlice []string) {
	innerSlice = append(innerSlice, "a")
	innerSlice[0] = "b"
	innerSlice[1] = "b"
	fmt.Println(innerSlice)
}

func main() {
	outerSlice := []string{"a", "a"}
	modifySlice(outerSlice)
	fmt.Print(outerSlice)
}
