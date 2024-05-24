package main

import "fmt"

const (
	PI = 3.14
)

type Circle struct {
	radius float64
}

func NewCircle(r float64) Circle {
	return Circle{radius: r}
}

func (c Circle) calculateCircumference() {
	circ := 2 * PI * c.radius
	fmt.Println(circ)
}

func main() {
	c := NewCircle(10)
	c.calculateCircumference()
}
