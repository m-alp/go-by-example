package examples

import "fmt"

// methods are defined on struct types
type rect struct {
	width, height int
}

func (r*rect) area() int { // receives a rect pointer
	return r.width * r.height
}

func (r rect) perim() int { // receives a rect instance
	return 2*r.width + 2*r.height
}

func (r rect) setHeight(h int) {
	r.height = h
}

func (r*rect) setHeightRef(h int) {
	r.height = h
}

func Methods() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

	fmt.Println("---")

	r2 := rect{}

	r2.setHeight(1)
	fmt.Println("h:", r2.height)

	r2.setHeightRef(1)
	fmt.Println("href:", r2.height)
}
