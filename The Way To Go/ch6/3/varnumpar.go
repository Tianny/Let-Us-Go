package main

import "fmt"

func main()  {
	x := min(1, 2, 3, 0)
	fmt.Printf("The minimum is %d\n", x)
	slice := []int{1,3,5,6}
	x = min(slice...)
	fmt.Printf("The minimum is %d\n", x)

	
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func typecheck(int ,int ,values...interface{}) {
	for _, value := range values {
		switch v := value.(type) {
			case int: …
			case float: …
			case string: …
			case bool: …
			default: …
		}
	}
}