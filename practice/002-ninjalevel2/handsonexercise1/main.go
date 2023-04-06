package main

import "fmt"

// write a program that prints a number in decimal, binary, and hex

func main() {
	n := 42

	fmt.Printf("%d\t%b\t%#x", n, n, n)

}
