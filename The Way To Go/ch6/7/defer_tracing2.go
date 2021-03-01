package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string)  {
	fmt.Println("leaving:", s)
}

func a()  {
	defer un(trace("a")) // 当执行到该条语句时，函数和参数表达式先得到计算，即 trace("a") 先可以获得计算
	fmt.Println("in a")
}

func b()  {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main()  {
	b()
}


// entering: b
// in b
// entering: a
// in a
// leaving: a
// leaving: b