package main

import "fmt"

func main() {
	var foo string = "Hello"

	// Pointer
	var bar *string = &foo

	fmt.Println("Variable", foo, &foo)
	fmt.Println("Pointer", bar, &bar)

	// Dereferencing
	fmt.Println("Dereferencing (bar)", *bar)

	// Modify pointer value
	*bar = "Bye"
	fmt.Println("Foo", foo)
}
