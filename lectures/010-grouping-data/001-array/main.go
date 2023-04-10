package main

import "fmt"

func main() {
	var x [5]int //array of int, size 5
	fmt.Println(x)

	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
}

// remember 0 based index
// GO is about ease of programming
