package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type traingle struct {
	height float64
	base   float64
}

func main() {
	s := square{sideLength: 4}
	t := traingle{height: 6, base: 4}
	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t traingle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println("Area = ", s.getArea())
}
