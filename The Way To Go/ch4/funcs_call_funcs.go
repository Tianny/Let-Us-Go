package main

var aFuncs string

func main() {
	aFuncs = "G"
	print(aFuncs)
	f1()
}

func f1() {
	aFuncs := "0"
	print(aFuncs)
	f2()
}

func f2() {
	print(aFuncs)
}
