package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

// keyword identifier type
// a value can be more than one type
func (s secretAgent) walk() {
	fmt.Println(s.first, "is walking")
}

func (p person) walk() {
	fmt.Println(p.first, "is walking")
}

type human interface {
	speak()
	walk()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrrrrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr. ",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
	p1.speak()

	sa1.walk()
	sa2.walk()
	p1.walk()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

	/*
		// conversion
		var x hotdog = 42
		fmt.Println(x)
		fmt.Printf("%T\n", x)
	*/
}
