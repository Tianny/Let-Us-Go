package main

import "fmt"

type B struct {
	thing int
}

func (b *B) change() {
	b.thing = 1
}

// 值拷贝
func (b B) write() string {
	b.thing = 2
	return fmt.Sprint(b)
}

func main() {
	var b1 B               // b1 是值
	b1.change()            // 修改了 b1  的值，b1=1
	fmt.Printf(b1.write()) // 只是修改了 b1 复制品的值，b1 的值未变
	fmt.Printf("%v", b1)   // b1 的值未变，b1=1

	fmt.Println("")

	b2 := new(B)           // b2 是指针
	b2.change()            // 修改了 b2 的值，b2=1
	fmt.Printf(b2.write()) // 只是修改了 b2 复制品的值，b2 的值未变，
	fmt.Printf("%v", b2)   // b2 的值未变，b2=1

}
