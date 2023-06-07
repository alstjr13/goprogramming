package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}

func foo(x ...int) int {
	var sum int

	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	var sum int

	for _, v := range x {
		sum += v
	}
	return sum
}
