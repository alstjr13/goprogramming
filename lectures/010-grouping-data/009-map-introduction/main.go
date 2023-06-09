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

	if v, ok := m["Miss Moneypenny"]; ok { // this is the comma ok idiom
		fmt.Println("THIS IS THE IF PRINT", v)
	}
}

// maps: key, value store
// FAST
// un-ordered list
