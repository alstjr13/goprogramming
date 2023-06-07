package main

import "fmt"

func main() {
	foo(1, 2, 3, 4, 5, 6, 7, 8, 9)
	li := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	bar(li)

}

func foo(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	fmt.Println("the total is: ", sum)
	return sum

}

func bar(ii []int) int {
	var sum int

	for _, v := range ii {
		sum += v
	}
	fmt.Println("the total bar is : ", sum)
	return sum

}
