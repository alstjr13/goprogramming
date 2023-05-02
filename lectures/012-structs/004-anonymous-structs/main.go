package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	/*
		p1 := person{
			first: "James",
			last:  "Bond",
			age:   42,
		}
	*/

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   42,
	}

	fmt.Println(p1)
}
