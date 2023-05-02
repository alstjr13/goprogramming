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

	fmt.Println(p1)
	fmt.Println(p2)

	for i, v := range p2.iceFav {
		fmt.Println(i, v)
	}

	fmt.Println(p1.first, "likes", p1.iceFav)
}

/* Create your own type "person" which will have an underlying type of "struct" so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two values of type person. Print out the values, ranging over the elements in the slice
*/
