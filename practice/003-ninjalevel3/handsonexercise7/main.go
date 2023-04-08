package main

import "fmt"

func main() {
	x := "a Bond"

	if x == "Money Penny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BONDBONDBOND")
	} else {
		fmt.Println("neither")
	}
}
