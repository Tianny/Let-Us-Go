package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer func ()  {
			fmt.Print(i) // 传递 i 的引用
		}()
	}

	for i := 0; i < 3; i++ {
		defer fmt.Print(i) // 传递 i 的拷贝
	}

}