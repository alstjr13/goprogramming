package main

import "fmt"

func main() {
	xs := []string{"James", "Bond", "Shaken, not stirred"}
	ys := []string{"Miss", "Moneypenny", "Helloooo, James."}

	ms := [][]string{xs, ys}

	fmt.Println(xs)
	fmt.Println(ys)

	fmt.Println(ms)

	for i, xs := range ms {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}
