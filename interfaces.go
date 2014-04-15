package main

import "fmt"

// To implement an interface in Go,
// we just need to implement all the methods in the interface.
type geometry interface {
	area() int
	perim() int
}

type rect struct {
	width, height int
}

type square struct {
	side int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func (s square) area() int {
	return s.side * s.side
}

func (s square) perim() int {
	return 4 * s.side
}

func measure(g geometry) {
	fmt.Println("Geometry:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perim())
}

func main() {
	r := rect{20, 30}
	s := square{25}

	measure(r)
	measure(s)
}
