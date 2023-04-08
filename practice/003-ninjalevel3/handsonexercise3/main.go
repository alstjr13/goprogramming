package main

import "fmt"

func main() {

	/*
		for i := 1998; i <= 2023; i++ {
			if i == 1998 {
				fmt.Printf("I have been born on %d. \n", i)
			} else {
				fmt.Printf("I have lived in %d. \n", i)
			}
		}
	*/

	birthYear := 1998

	for birthYear <= 2023 {
		fmt.Println(birthYear)
		birthYear++
	}
}
