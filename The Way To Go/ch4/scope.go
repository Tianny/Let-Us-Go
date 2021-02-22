package main

var aScope = "G"

func main() {
	n()
	m()
	n()
}

func m() {
	// a_scope:= "0"
	aScope = "0"
	print(aScope)
}

func n() { print(aScope) }
