package main

// 函数作为返回值
func setup(task string) func() {
	println("do some setup stuff for", task)
	return func() {
		println("do some teardown stuff for", task)
	}
}

func main() {
	teardown := setup("demo")
	defer teardown()
	println("do some business stuff")
}
