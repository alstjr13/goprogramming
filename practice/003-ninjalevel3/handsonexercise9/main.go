package main

import "fmt"

func main() {
	favSport := "rugby"
	switch favSport {
	case "skiing":
		fmt.Println("go to mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to the ocean!")
	case "rugby":
		fmt.Println("go to the UK!")
	}
}
