package main

import "sync"

func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}

	// 启动一个监控的 goroutine，用于等待所有的的执行 goroutine 退出
	go func() {
		wg.Wait()
		close(out)
	}()

	var result []int
	for v := range out {
		result = append(result, v)
	}

	return result
}
