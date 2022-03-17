package main

import "fmt"

func main() {
	src := []int{1,2,3}
	dst := make([]*int, 3, 3)
	for idx, val := range src {
		dst[idx] = &val
	}

	fmt.Print("dst value:")
	for _, val := range dst {
		fmt.Print(" ", *val)
	}
}
