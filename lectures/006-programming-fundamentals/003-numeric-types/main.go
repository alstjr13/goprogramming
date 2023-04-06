package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64

var z int8

var a byte //same as uint8
var b rune //same as int32

func main() {
	x = 42
	y = 42.34534
	z = -127

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(runtime.GOOS)   // windows
	fmt.Println(runtime.GOARCH) // amd64

}

// Floating point numbers have decimals
// Integers don't have decimals
