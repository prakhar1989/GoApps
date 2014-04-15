package main

import "fmt"

type person struct {
	name string
	age  int
}

type rect struct {
	width, height int
}

// defining a method on the struct
func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Fred", age: 20}) // using maps
	fmt.Println(person{name: "Fred"})          // missed out fields are zero

	fmt.Println(&person{name: "Fred", age: 20}) // pointer to a struct

	s := person{name: "Prakhar", age: 24}
	fmt.Println(s.name, "is", s.age, "years old")

	sp := &s    // pointer to the struct
	sp.age = 40 // structs are mutable
	fmt.Println(s.name, "is now", s.age, "years old")

	r := rect{width: 40, height: 20}
	fmt.Println("Area of the rect", r.area())
	fmt.Println("Perimeter of the rect", r.perim())
}
