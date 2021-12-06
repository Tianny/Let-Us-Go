package main

import (
	"fmt"
	"math"
)

type square struct {
	side float32
}

type circle struct {
	radius float32
}

type shaper interface {
	area() float32
}

func (sq *square) area() float32 {
	return sq.side * sq.side
}

func (ci *circle) area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {

	var areaIntf shaper

	sq1 := new(square)
	sq1.side = 5

	areaIntf = sq1
	if t, ok := areaIntf.(*square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	}

}

// 如果忽略 areaIntf.(*Square) 中的 * 号，会导致编译错误：impossible type assertion: Square does not implement Shaper (Area method has pointer receiver)。
