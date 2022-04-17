package main

import "fmt"

type Human struct {
	Name string
}

func main() {
	foo := Human{
		Name: "Pedro",
	}

	foo.greet()
	foo.sayGoodbye()
}

func (h Human) greet() {
	fmt.Println("Hello", h.Name)
}

func (h Human) sayGoodbye() {
	fmt.Println("Goodbye", h.Name)
}
