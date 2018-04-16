package examples

import "fmt"

// sums two ints
func plus(a int, b int) int {
	return a + b
}

// sums three ints
func plusPlus(a, b, c int) int { // if multiple params with same type, only last needs type
	return a + b + c
}

// multiple return values
func vals() (int, int) {
	return 3, 7
}

// variadic function (can accept any number of args)
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// anonymous function (clojure)
func intSeq() func() int  {
	i := 0
	return func() int {
		i++
		return i
	}
}

// recursive function
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func Functions() {
	// basic functions
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	fmt.Println("---")

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	fmt.Println("---")

	// variadic function
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4} // slice
	sum(nums...) // -> func(slice...)

	fmt.Println("---")

	// anonymous function (clojure)
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())

	fmt.Println("---")

	// recursive function
	fmt.Println(fact(7))

}