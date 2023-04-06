package main

import "fmt"

const (
	a = iota + 2020
	b = iota + 2020
	c = iota + 2020
	d = iota + 2020
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
