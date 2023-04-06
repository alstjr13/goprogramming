package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c) // 2

	fmt.Println(d) // 0
	fmt.Println(e) // 1
	fmt.Println(f) // 2
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
