package main

import "fmt"

func main() {
	fmt.Println("Hello, Playground")

	f := func() {
		fmt.Println("my first expression")
	}
	f()

	g := func(x int) {
		fmt.Println("the year big brother started watching:", x)
	}
	g(1984)
}
