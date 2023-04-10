package main

import "fmt"

func main() {
	// x := type{values}  // COMPOSITE LITERAL

	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println(len(x))
}

//slicing a slice

//Composite data type (Compound data type):
// any data type which can be constructed in a program using the programming language's primitive data types and other composite types.
// aka structure or aggregate data type
