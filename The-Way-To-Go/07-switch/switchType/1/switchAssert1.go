package main

func main() {
	var x interface{} = 13

	switch x.(type) {
	case nil:
		println("x is nil")
	case int:
		println("x is int")
	default:
		println("doesn't support the type")
	}
}
