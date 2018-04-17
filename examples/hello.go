package examples

import "fmt"

func Hello() {
	fmt.Print("Hello ", "World", "!", "\n")

	fmt.Println("Hello World!")

	fmt.Printf("%s %s!\n", "Hello", "World")

	a := fmt.Sprintf("%s!", "Hello World")
	fmt.Println(a)
}