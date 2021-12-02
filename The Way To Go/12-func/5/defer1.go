package main

import "fmt"

func main() {
	a() // 打印 0

}

func a() {
	i := 0
	defer fmt.Println(i) // 输出 0
	i++
	return
}
