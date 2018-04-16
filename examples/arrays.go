package examples

import "fmt"

// arrays always have a type, all its elements must be of that type
func Arrays() {
	var a [5]int // array with 5 ints
	fmt.Println("emp:", a) // empty

	a[4] = 100 // set
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a)) // length

	fmt.Println("---")

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	fmt.Println("---")

	c := [...]string{"Penn", "Teller"} // compiler counts the array elements for us
	fmt.Println("str:", c)

	fmt.Println("---")

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}