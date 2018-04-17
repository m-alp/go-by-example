package examples

import "fmt"

// arguments passed by value
func zeroval(ival int) { // receives a copy of ival, not the original
	ival = 0 // sets the value to 0
}

// reference is passed
func zeroptr(iptr *int) {
	*iptr = 0 // dereferences the int pointer from memory address
}

func Pointers()  {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
