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

	// 变量 t 得到了 areaIntf 的值和类型
	switch t := areaIntf.(type) {
	case *square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

	// 如果仅仅是测试变量的类型，不用它的值，那么就可以不需要赋值语句
	switch areaIntf.(type) {
	case *square:
		// TODO
	case *circle:
		// TODO
	case nil:
		// TODO
	default:
		// TODO
	}
}
