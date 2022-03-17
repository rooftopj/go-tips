package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := make([]int32, 1, 10)
	s1[0] = 1
	fmt.Printf("s1 pointer: %p\n", s1)
	fmt.Println("s1 value before:", s1)
	s2 := append(s1, 2)
	fmt.Printf("s2 pointer: %p\n", s2)
	fmt.Println("s2 value:", s2)
	fmt.Println("s1 value after:", s1)
	fmt.Println("len of s1:", len(s1))
	fmt.Println("len of s2:", len(s2))
	fmt.Println("s1[1] value:", *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&s1[0])) + unsafe.Sizeof(int32(0)))))
}