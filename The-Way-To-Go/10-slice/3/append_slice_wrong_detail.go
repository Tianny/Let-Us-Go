package main

import "fmt"

func modifySliceDetail(innerSlice []string) {
	fmt.Printf("%p %v   %p\n", &innerSlice, innerSlice, innerSlice)
	innerSlice = append(innerSlice, "a")
	innerSlice[0] = "b"
	innerSlice[1] = "b"
	fmt.Printf("%p %v %p\n", &innerSlice, innerSlice, innerSlice)
}

func main() {
	outerSlice := []string{"a", "a"}
	fmt.Printf("%p %v   %p\n", &outerSlice, outerSlice, outerSlice)
	modifySliceDetail(outerSlice)
	fmt.Printf("%p %v   %p\n", &outerSlice, outerSlice, outerSlice)
}
