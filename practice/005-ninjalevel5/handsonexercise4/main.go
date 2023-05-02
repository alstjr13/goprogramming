package main

import "fmt"

func main() {
	p1 := struct {
		first      string
		last       string
		age        int
		isHandsome bool
	}{
		first:      "James",
		last:       "Bond",
		age:        43,
		isHandsome: true,
	}

	fmt.Println(p1)
	fmt.Println(p1.first)
}
