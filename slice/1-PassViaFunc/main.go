package main

import "fmt"

func main() {
	s := []int{1,2,3,4,5}
	fmt.Printf("before pointer: %p\n", &s)
	fmt.Println("before:", s)
	changeSlice(s)
	fmt.Println("after:", s)
}

func changeSlice(s []int) {
	if len(s) <= 0 {
		return
	}
	s[0] = s[0] + 10
	fmt.Printf("after pointer: %p\n", &s)
}