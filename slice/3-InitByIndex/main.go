package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := s1[0:3]
	fmt.Println("before s1: ", s1)
	s2[0] = 11
	fmt.Println("after revise, s1: ", s1)
	s2 = append(s2, 14)
	fmt.Println("after append, s1: ", s1)
}
