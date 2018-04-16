package examples

import "fmt"

func Maps() {
	m := make(map[string]int) // -> make(map[key-type]val-type)
	fmt.Println("emp:", m) // empty

	// set
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// get
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m)) // length

	fmt.Println("---")

	delete(m, "k2") // delete key/val
	fmt.Println("map:", m)

	_, prs := m["k2"] // key exists. First return is get, second is bool for exists.
	fmt.Println("prs:", prs)

	fmt.Println("---")

	n := map[string]int{"foo": 1, "bar": 2} // other way to declare/init
	fmt.Println("map:", n)
}
