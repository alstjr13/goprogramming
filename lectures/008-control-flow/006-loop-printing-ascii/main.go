package main

import "fmt"

//33 ~ 122
func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}
}
