package main

import (
	"fmt"
	"math"
)

// Any type that has the two methods LanguageName() and Greet(name string) automatically fits the Greeter interface.

/* INTERFACE
Interfaces in Go define a set of method signatures that types must implement to satisfy the interface.
- Type implemenets an interface if it has all the methods required by that interface. No explicit declaration is needed.
type InterfaceName interface {
    MethodName(parameters) returnType
}

- Emmpty interface (interface{}) can be used to represent any type, as all types implement it
- Type asssertions can be used to exract the underlying type from an interface
value := interfaceVariable.(ConcreteType)
- Type switch allows branching logic based on the concrete type stored in an interface
switch v := i.(type) {
    case string:
        // handle string type
}
- Zero value of an interface is 'nil' - no concrete value or type is stored
- Functions can accept interfaces as parameters, allowing flexible and polymorphic behavior
- Interfaces can embed other interfaces to create more complex interfaces (compositions)

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
