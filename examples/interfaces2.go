package examples

import "fmt"

// we’ll define an Animal as being anything that can speak
// Animal: any type with method named Speak
type Animal interface {
	Speak() string
}

// let’s create a four types that satisfy this interface
type Dog struct {}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

type Llama struct{}

func (l Llama) Speak() string {
	return "????"
}

type JavaProgrammer struct {}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func Interfaces2()  {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
