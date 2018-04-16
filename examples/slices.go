package examples

import "fmt"

func Slices() {
	s := make([]string, 3) // -> func make([]T, len, cap = optional) []T
	fmt.Println("emp:", s) // empty

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s)) // length

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	fmt.Println("---")

	c := make([]string, len(s))
	copy(c, s) // copy array from low to high (if len is smaller, last elements are not copied)
	fmt.Println("cpy:", c)

	fmt.Println("---")

	l := s[2:5] // slice[low:high]
	fmt.Println("sl1", l)

	l = s[:5] // slice up to [5] (excluding)
	fmt.Println("sl2", l)

	l = s[2:] // slice up from [2] (including)
	fmt.Println("sl3", l)

	fmt.Println("---")

	t := []string{"g", "h", "i"} // like array, without count in the beginning
	fmt.Println("dcl:", t)

	fmt.Println("---")

	multiDim()
}

func multiDim() {
	twoD := make([][]int, 3)

	fmt.Println(twoD)

	for i:= 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
