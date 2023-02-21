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
	for i := range countTo(10) {
		// 提前退出循环，goroutine 将会永远阻塞，等待从通道中读取一个值
		// 接收者发生退出，停止继续接收上游数据，发送者就会被阻塞
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
}
