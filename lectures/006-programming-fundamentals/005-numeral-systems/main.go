package main

import "fmt"

func main() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#X\n", n)

}

// 42 in binary
// 101010

// base 16
// A B C D E F
// 0 1 2 3 4 5 6 7 8 9

// 911 in hexadecimal
// 38F
