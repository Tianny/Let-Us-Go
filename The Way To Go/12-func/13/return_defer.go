package main

import "fmt"

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
func main() {
	fmt.Println(f())
}

// ret++ 是在执行 return 1 语句后发生的
// 2
