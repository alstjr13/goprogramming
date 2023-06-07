package main

import "fmt"

func main() {
	defer foo()
	defer bar()
	zoo()
	defer bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("I'm in bar")
}

func zoo() {
	fmt.Println("I'm in zoo")
}

// deferred functions run in LIFO order
// last in first out : LIFO
