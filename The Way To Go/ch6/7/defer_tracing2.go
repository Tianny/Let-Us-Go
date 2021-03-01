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
	// 当执行到 defer 修饰的语句时，函数和参数表达式均可得到计算，但直到包含该 defer 语句的函数执行完毕时，defer 后的函数才会被执行
	// 所以这里 trace("a") 会按正常顺序执行
	defer un(trace("a")) 
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