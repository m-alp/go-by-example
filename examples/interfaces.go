package examples

import (
	"math"
	"fmt"
)

// an interface is a) a set of methods and b) a type
type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64  {
	return r.width * r.height
}
func (r rectangle) perim() float64  {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64  {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64  {
	return 2 * math.Pi * c.radius
}

func measure(g geometry)  { // g has interface type, so g can call implemented methods
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func Interfaces()  {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	// r & c both impl geometry, so we can use instances of these structs as arguments of measure()
	measure(r)
	measure(c)
}