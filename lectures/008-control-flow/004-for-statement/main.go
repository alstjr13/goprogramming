package main

import "fmt"

// "for" is a keyword

func main() {
	/*x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done")*/

	// init, condition, post
	/*
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		fmt.Println("done!")
	*/

	x := 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")
}
