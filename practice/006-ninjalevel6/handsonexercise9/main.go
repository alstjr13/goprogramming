package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("this is an anonymous func")
	}()

	x()

}

func foo() {
	fmt.Println("this is foo")
}

var x = func() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
