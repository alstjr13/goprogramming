package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("hi my name is", p.first, "and I am ", p.age, "years old")
}

func main() {
	p1 := person{
		first: "Ian",
		age:   29,
	}

	p1.speak()

}
