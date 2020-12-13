package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	sqr := square{
		sideLength: 1.0,
	}

	printArea(sqr)

	tri := triangle{
		height: 1.0,
		base:   1.0,
	}

	printArea(tri)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.height * t.base / 2.0
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
