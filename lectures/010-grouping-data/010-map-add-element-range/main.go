package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)

	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	m["todd"] = 33 // this is how you add stuff to the map

	if v, ok := m["Miss Moneypenny"]; ok { // this is the comma ok idiom
		fmt.Println("THIS IS THE IF PRINT", v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 6, 7, 8, 9, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}

// maps: key, value store
// FAST
// un-ordered list
