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

	var areaInter shaper
	sq1 := new(square)
	sq1.side = 5
	areaInter = sq1

	// 变量 t 得到了 areaInter 的值和类型
	switch t := areaInter.(type) {
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
	switch areaInter.(type) {
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
