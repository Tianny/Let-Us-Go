package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time // anonymous field
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]

}

func main() {
	m := myTime{time.Now()}
	// 调用匿名 Time 上的 String方法
	fmt.Println("Full time now:", m.Time.String())
	// 调用 myTime.first3Chars
	fmt.Println("First 3 chars:", m.first3Chars())
}
