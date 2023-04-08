package main

import "fmt"

func main() {

	/*
		x := 41 / 40 // 1
		y := 83 % 40 // 3
		fmt.Println(x)
		fmt.Println(y)
	*/

	// print out even numbers
	x := 0
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)

	}

}

// % modulo operation
