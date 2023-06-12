package main

import "fmt"

var hotdog int

type person struct {
}
type human interface {
	speak()
}

func main() {
	n := foo()
	x, s := bar()

	fmt.Println(n, x, s)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 1984, "Big Brother is watching"
}

// func, receiver, identifier, params, returns
