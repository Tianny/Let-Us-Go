package main

import "fmt"

type shaper interface {
	area() float32
}

type square struct {
	side float32
}

func (sq *square) area() float32 {
	return sq.side * sq.side
}

type rectangle struct {
	length, width float32
}

func (r rectangle) area() float32 {
	return r.length * r.width
}

func main() {
	r := rectangle{5,3}
	q := &square{5}
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []shaper{r,q}
	fmt.Println("Looping through shapes for area ...")
	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].area())
	}

}