package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	//fmt.Println(x)

	x["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	for k, v := range x {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i+1, v2)
		}
	}

	delete(x, "no_dr")

	for k, v := range x {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i+1, v2)
		}
	}

}
