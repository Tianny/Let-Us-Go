package main

import "fmt"

// 每个 goroutine 的闭包都捕获了同一个变量

func main() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func() {
			ch <- v * 2
		}()
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
