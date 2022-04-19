package main

import "fmt"

func main() {
	var arr1 [5]int // 类型是 [5]int
	fmt.Println(arr1)

	var arr2 = new([5]int) // 类型是指针类型 *[5]int
	// Go 中数组是值类型，所以可以通过 new() 来创建
	fmt.Println(arr2)

	a := [10]int{}
	a[0] = 2
	fmt.Println(a)

	p := new([10]int)
	p[0] = 1 // 指向数组的指针也可以直接使用索引形式
	fmt.Println(p)

	a1 := [...]string{"a", "b", "c", "d"}
	for i := range a1 {
		fmt.Println("Array item", i, "is", a1[i])
	}

	var arr = [5]int{}
	fmt.Println(arr)

	var arrAge = [5]int{18, 20, 15, 22, 16}
	fmt.Println(arrAge)

	var arrLazy = [...]int{5, 6, 7, 8, 22}
	fmt.Println(arrLazy)

	arr3 := [10]int{1, 2, 3} // 这是一个有 10 个元素的数组，除了前三个元素外其他元素都为 0
	fmt.Println(arr3)

	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"} // 只有索引 3 和 4 被赋予实际的值，其他元素都被设置为空的字符串
	fmt.Println(arrKeyValue)

	r := [...]int{99: -1} // 定义了一个含有100个元素的数组r，最后一个元素被初始化为 -1，其它元素都是用0初始化
	fmt.Println(r)

	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a { // dereferencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}
