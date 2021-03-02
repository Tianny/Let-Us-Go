package main

import (
	"fmt"
	"strings"
)

func main() {
	addBmp := makeAddSuffix(".bmp")
	addJpeg := makeAddSuffix(".jpeg")
	fmt.Println(addBmp("file")) // returns: file.bmp
	fmt.Println(addJpeg("file"))  // returns: file.jpeg
}

func makeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}