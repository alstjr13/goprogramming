package main

import "fmt"

type person struct {
	first  string
	last   string
	iceFav []string
}

func main() {
	p1 := person{
		first:  "James",
		last:   "Bond",
		iceFav: []string{"chocolate", "strawberry", "vanila"},
	}

	p2 := person{
		first:  "Ed",
		last:   "Sheeran",
		iceFav: []string{"chocolate", "almond", "lemon"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)

		for i, val := range v.iceFav {
			fmt.Println(i, val)
		}
		fmt.Println("------------")
	}
	/*
		fmt.Println(p1)
		fmt.Println(p2)

		for i, v := range p2.iceFav {
			fmt.Println(i, v)
		}

		fmt.Println(p1.first, "likes", p1.iceFav)
	*/
}
