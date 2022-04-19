package main

import "fmt"

func main() {
	p := new([]int) // *p == nil; with len and cap 0
	fmt.Println(p)

	var v []int = make([]int, 10, 50) // 切片已经初始化，且关联数组已经创建，不过是一个空数组
	fmt.Println(v)
}
