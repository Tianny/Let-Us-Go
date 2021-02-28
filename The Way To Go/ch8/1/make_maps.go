package main

import "fmt"

func main() {
	mapLit := map[string]int{"one": 1,"two": 2}
	fmt.Println(mapLit)

	mapCreated := make(map[string]float32)
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159

	fmt.Println(mapCreated)

}