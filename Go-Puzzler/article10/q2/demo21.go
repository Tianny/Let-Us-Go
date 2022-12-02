package main

func main() {
	// 示例 1
	ch1 := make(chan int, 1)
	ch1 <- 1
	// ch1 <-2 // 通道已满，因此这里会造成阻塞

	// 示例 2
	ch2 := make(chan int, 1)
	elem, ok := <-ch2 // 通道已空，因此这里会造成阻塞
	_, _ = elem, ok
	ch2 <- 1

	// 示例 3
	var ch3 chan int
	// ch3 <- 1 // 通道的值为 nil，这里会造成永久的阻塞
	// <-ch3 // 通道值为 nil，这里会造成永久的阻塞
	_ = ch3
}
