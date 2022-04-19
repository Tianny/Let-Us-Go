package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("子 go 程结束")

		fmt.Println("子 go 程正在运行……")

		c <- 666 // 666 发送到 c
	}()
	fmt.Println(1)

	num := <-c // 从 c 中接收数据，并赋值给 num

	fmt.Println("num = ", num)
	fmt.Println("main go 程结束")
}
