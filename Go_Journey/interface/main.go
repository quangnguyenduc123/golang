/////////////////////////////////
// Implementing Interfaces in Go
// Go Playground: https://play.golang.org/p/SMjFrOYL5f3
/////////////////////////////////

package main

import (
	"fmt"
	"math"
)

// declaring an interface type called shape
// an interface contains only the signatures of the methods, but not their implementation
type shape interface {
	area() float64
	perimeter() float64
}

// declaring 2 struct types that represent geometrical shapes: rectangle and circle

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

// method that calculates circle's area
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// method that calculates rectangle's area
func (r rectangle) area() float64 {
	return r.height * r.width
}

// method that calculates circle's perimeter
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// method that calculates rectangle's perimeter
func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// declaring a method for circle type
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

// any type that implements the interface is also of type of the interface
// rectangle and circle values are also of type shape
func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func assertAndSwitch() {
	// declaring an interface value that holds a circle type value
	var s shape = circle{radius: 2.5}

	fmt.Printf("%T\n", s) //interface dynamic type is circle

	// no direct access to interface's dynamic values
	// s.volume() -> error

	// there is access only to the methods that are defined inside the interface
	fmt.Printf("Circle Area:%v\n", s.area())

	// an interface value hides its dynamic value.
	// use type assertion to extract and return the dynamic value of the interface value.
	fmt.Printf("Sphere Volume:%v\n", s.(circle).volume())

	// checking if the assertion succeded or not
	ball, ok := s.(circle)
	if ok {
		fmt.Printf("Ball Volume:%v\n", ball.volume())
	}

	//** TYPE SWITCHES **//

	// it permits several type assertions in series
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)

	}
}
func main() {
	var s shape

	//An interface type can contain a reference to an instance of any of the types
	//that implement the interface (an interface has what is called a dynamic type) change at runtime
	ball := circle{radius: 5}
	s = ball
	print(s)

	room := rectangle{width: 2, height: 3}
	s = room
	print(s)
}
