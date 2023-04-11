package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	y := []int{234, 456, 678, 987}

	x = append(x, y...) // [4,5,7,8,42,234,456,678,987]
	fmt.Println(x)

	x = append(x[:2], x[4:]...) // [4,5,42,234,456,678,987]
	fmt.Println(x)

}
