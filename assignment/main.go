package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	size float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {

	s := square{5}
	t := triangle{5, 1}

	printArea(s)
	printArea(t)

}

func (s square) getArea() float64 {
	return float64(s.size * s.size)
}

func (t triangle) getArea() float64 {
	return float64(0.5 * t.base * t.height)
}

func printArea(s shape) {
	fmt.Println("Area : ", s.getArea())
}
