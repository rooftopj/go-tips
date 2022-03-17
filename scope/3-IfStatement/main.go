package main

import "fmt"

func main() {
	if x := 1; false {
		fmt.Println("if, x:", x)
		y := 2
		fmt.Println("if, y:", y)
	} else if m := 3; x > 2 {
		fmt.Println("else if, m:", m)
		n := 4
		fmt.Println("else if, n:", n)
	} else {
		fmt.Printf("else, before, x: %v, m: %v\n", x, m)
		x = 11
		//y = 12 error: undefined: y
		m = 13
		//n = 14 error: undefined: n
		fmt.Printf("else, after, x: %v, m: %v\n", x, m)
	}
}
