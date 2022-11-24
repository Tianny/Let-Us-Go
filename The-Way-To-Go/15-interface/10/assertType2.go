package main

import "fmt"

type MyInterface interface {
	M1()
}

type T int

func (T) M1() {
	println("T's M1")
}

func main() {
	var t T
	var i interface{} = t
	v1, ok := i.(MyInterface)
	if !ok {
		panic("the value of i is not MyInterface")
	}
	v1.M1()
	fmt.Printf("the type of v1 is %T\n", v1)

	// 变量 i 必须是接口类型的变量。否则编译器会报错：invalid type assertion: i.(T) (non-interface type (type of i) on left)
	//j := int64(3)
	//v2, ok := j.(MyInterface)
	//fmt.Printf("the type of v2 is %T\n", v2)
}
