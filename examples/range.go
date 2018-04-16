package examples

import "fmt"

// range iterates over elements in some data structures
func Range() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { // only values
		sum += num
	}
	fmt.Println("sum:", sum)

	fmt.Println("---")

	for i, num := range nums { // keys and values
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	fmt.Println("---")

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s => %s\n", k, v)
	}

	fmt.Println("---")

	for k := range kvs { // only keys
		fmt.Println("key:", k)
	}

	fmt.Println("---")

	for i, c := range "go" { // range on strings iterates over Unicode code points
		fmt.Println(i, c)
	}
}
