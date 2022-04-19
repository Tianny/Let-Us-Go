package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func printer(str string) {
	mutex.Lock()         // 添加互斥锁
	defer mutex.Unlock() // 使用结束时解锁

	for _, data := range str { // 迭代器
		fmt.Printf("%c", data)
	}
	fmt.Println()
}

func person1(s1 string) {
	printer(s1)
}

func person2() {
	printer("world") // 调函数时传参
}

func main() {
	go person1("hello") // main 中传参
	go person2()
	select {}
}
