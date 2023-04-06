package main

import "fmt"

// Create TYPED and UNTYPED constants. Print the values of the constants

const (
	a     = 42
	b int = 42
)

func main() {
	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}
