package main

import "fmt"

func main() {
	x := [5]int{}

	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5

	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\t", x) // print out the type of the array
}
