package main

import "fmt"

func main() {
	x := make([]string, 7, 7)
	fmt.Println("First Time")
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, `Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`)

	fmt.Println("Second Time")
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// CORRECT CODE IS BELOW
	y := make([]string, 7, 7)
	fmt.Println("Third Time WITH Y")
	fmt.Println(len(y))
	fmt.Println(cap(y))

	states := []string{`Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`}

	for i, v := range states {
		y[i] = v
	}

	fmt.Println(("Fourth Time WITH Y"))
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	for i := 0; i < len(y); i++ {
		fmt.Println(i, y[i])
	}
}
