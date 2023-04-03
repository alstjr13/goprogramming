package main

import "fmt"

var x bool

func main() {
	a := 6
	b := 42
	fmt.Println(x) // false
	x = true

	fmt.Println(a != b) // true
	fmt.Println(x)      // true
}

// Boolean is either "true" or "false"
