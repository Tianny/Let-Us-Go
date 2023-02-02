package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d ", i)
	}
	wg.Done()
}

func main() {
	c := make(chan int)
	wg.Add(1)
	go printer(c)

	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)

	wg.Wait()
}
