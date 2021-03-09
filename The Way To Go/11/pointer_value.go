package main

import "fmt"

type B struct {
	thing int
}

func (b *B) change()  {
	b.thing = 1
}

func (b B) write() string {
	b.thing = 2
	return fmt.Sprint(b)
}

func main() {
	var b1 B // b1 是值
	b1.change()
	fmt.Printf(b1.write())
	fmt.Printf("%v" ,b1)

	fmt.Println("")

	b2 := new(B) // b2 是指针
	b2.change()
	fmt.Printf(b2.write())
	fmt.Printf("%v" ,b2)
}