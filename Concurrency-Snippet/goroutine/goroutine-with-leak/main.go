package main

import "fmt"

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	// 正常情况下，当遍历完所有的值时，goroutine 将会自动退出
	for i := range countTo(10) {
		fmt.Println(i)
	}
}
