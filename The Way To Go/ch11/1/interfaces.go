package main

import (

	"fmt"
)

type shaper interface {
	area() float32
}

type square struct {
	side float32
}

func (sq *square) area()  float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(square)
	sq1.side = 5
	fmt.Printf("The square has area: %f\n", sq1.area())

	var areaIntf shaper
	// 将一个 Square 类型的变量赋值给一个接口类型的变量
	// 现在接口变量包含一个指向 Square 变量的引用，通过它可以调用 Square 上的方法 Area()
	// 这就是多态的 Go 版本
	areaIntf = sq1 
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.area())

}