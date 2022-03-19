package main

import "fmt"

type S struct {
	a int
	b int
	t string
}

func (s *S) pointer_method() {
	fmt.Printf("%s, pointer_method, address: %p\n", s.t, s)
	s.a = s.a + 10
}

func (s S) value_method() {
	fmt.Printf("%s, value_method, address: %p\n", s.t, &s)
	s.b = s.b + 10
}

func main() {
	s_value := S{1, 2, "value entity"}
	s_pointer := &S{3, 4, "pointer entity"}
	fmt.Printf("s_value's address: %p\n", &s_value)
	fmt.Printf("s_pointer's address: %p\n", s_pointer)
	s_value.pointer_method()
	s_value.value_method()
	s_pointer.pointer_method()
	s_pointer.value_method()
	fmt.Println("s_value:", s_value)
	fmt.Println("s_pointer:", s_pointer)
}
