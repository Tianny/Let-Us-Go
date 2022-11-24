package main

func main() {
	var x interface{} = 13
	switch v := x.(type) {
	case int:
		println("the type of v is int, v=", v)
	default:
		println("no match type")
	}
}
