package main

func main() {

	const Pi = 3.14

	const b string = "go"

	const b11 = "go" // 类型推断

	const Monday, Tuesday, Wednesday = 1, 2, 3 // 并行赋值

	// 枚举使用
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	// iota
	const (
		apple  = iota // 0
		banana = iota // 1
		cat    = iota // 2
	)

	// 赋值一个常量时，之后没赋值的常量都会应用上一行的赋值表达式
	const (
		a1 = iota // a1 = 0
		b1        // b1 = 1
		c1        // c1 = 2
		d1 = 5    // d1 = 5
		e1        // e1 = 5
	)

	// 一次性定义两个常量，只算一行，两个iota 表达式的值一样
	// 赋值两个常量，iota 只会增长一次，而不会因为使用了两次就增长两次
	const (
		Apple, Banana   = iota + 1, iota + 2 // Apple=1 Banana=2
		Elderberry, Fig                      // Elderberry=2, Fig=3
	)

	const (
		_  = iota             // 使用 _ 忽略不需要的 iota
		KB = 1 << (10 * iota) // 1 << (10*1)
		MB                    // 1 << (10*2)
		GB                    // 1 << (10*3)
		TB                    // 1 << (10*4)
		PB                    // 1 << (10*5)
		EB                    // 1 << (10*6)
		ZB                    // 1 << (10*7)
		YB                    // 1 << (10*8)
	)
}
