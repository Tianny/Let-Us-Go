package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var count int           // 全局变量count
var rwlock sync.RWMutex // 全局读写锁 rwlock

func read(n int) {
	rwlock.RLock()
	fmt.Printf("读 goroutine %d 正在读取数据...\n", n)
	num := count
	fmt.Printf("读 goroutine %d 读取数据结束，读到 %d\n", n, num)
	defer rwlock.RUnlock()
}
func write(n int) {
	rwlock.Lock()
	fmt.Printf("写 goroutine %d 正在写数据...\n", n)
	num := rand.Intn(1000)
	count = num
	fmt.Printf("写 goroutine %d 写数据结束，写入新值 %d\n", n, num)
	defer rwlock.Unlock()
}

func main() {
	for i := 0; i < 5; i++ {
		go read(i + 1)
	}
	for i := 0; i < 5; i++ {
		go write(i + 1)
	}
	select {}
}
