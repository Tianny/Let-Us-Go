package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 示例 1
	// 只能发不能收的通道
	var uselessChan = make(chan<- int, 1)
	// 只能收不能发的通道
	var anotherUselessChan = make(<-chan int, 1)
	// 这里打印的是可以分别代表两个通道的指针的十六进制表示
	fmt.Printf("The useless channels: %v, %v\n", uselessChan, anotherUselessChan)

	// 示例 2
	intChan1 := make(chan int, 3)
	SendInt(intChan1)

	// 示例 4
	intChan2 := getIntChan()
	for elem := range intChan2 {
		fmt.Printf("The element in intChan2: %v\n", elem)
	}

	// 示例 5
	_ = GetIntChan(getIntChan)
}

// SendInt 示例 2
func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000)
}

// Notifier 示例 3
type Notifier interface {
	SendInt(ch chan<- int)
}

// 示例 4
func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

// GetIntChan 示例 5
type GetIntChan func() <-chan int
