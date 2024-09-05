package main

import (
	"fmt"
	"math"
)

/* INTERFACE
Interfaces are named collections of method signatures.

https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
*/

type geometry interface { // basic interface for geometric shapes
	area() float64
	perim() float64
}

type rect_i struct { // we create rect and circle types for our example
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect_i) area() float64 { // to implement an interface in Go we just need to implement all the methods in the interface
	return r.width * r.height
}
func (r rect_i) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 { // implementation for circles
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) { // if a variable has an interface type, then we can call methods that are in the named interface
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func interfaces() {
	r := rect_i{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r) // the circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure
	// {3 4}
	// 12
	// 14

	measure(c)
	// {5}
	// 78.53981633974483
	// 31.41592653589793

}
