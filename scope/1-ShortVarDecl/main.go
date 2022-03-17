package main

import (
	"fmt"
)

func main() {
	var a = 1
	b := 1
	{
		a := 2
		b := 2
		fmt.Printf("in the scope declare, a: %v, b: %v\n", a, b)
		a = 3
		b = 3
		fmt.Printf("in the scope cover, a: %v, b: %v\n", a, b)
	}
	fmt.Printf("out of the scope, a: %v, b: %v\n", a, b)
}
