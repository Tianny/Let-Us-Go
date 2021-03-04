package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Print(i) // 传递 i 的引用
		}()
	}

	for i := 0; i < 3; i++ {
		defer fmt.Print(i) // 在定义时已经获得了 i 的拷贝
	}

}

// 如果函数体内某个变量作为 defer 时匿名函数的参数，则在定义 defer 时即已经获得了拷贝，否则则是引用某个变量的地址
