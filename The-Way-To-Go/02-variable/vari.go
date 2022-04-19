package main

import "fmt"

func main() {
	var a = 15
	b := 15                // 常用
	a, b, c := 5, 7, "abc" // 并行赋值
	fmt.Println(a, b, c)

	_, b = 5, 7 // 值 5 被抛弃
}
