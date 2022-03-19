package main

import "fmt"

type S struct {
	a int
	b int
	t string
}

func (s *S) pointer_method() {
	fmt.Printf("%s, pointer_method, a: %d\n", s.t, s.a)
}

func (s S) value_method() {
	fmt.Printf("%s, value_method, b: %d\n", s.t, s.b)
}

func get_pointer_enetity() *S {
	return &S{3, 4, "pointer entity"}
}

func get_value_entity() S {
	return S{1, 2, "value entity"}
}

func main() {
	s_value := S{1, 2, "value entity"}
	s_pointer := &S{3, 4, "pointer entity"}
	s_value.pointer_method()
	s_value.value_method()
	s_pointer.pointer_method()
	s_pointer.value_method()

	// cannot call pointer method on get_value_entity()
	// cannot take the address of get_value_entity()
	// get_value_entity().pointer_method()

	get_value_entity().value_method()
	get_pointer_enetity().pointer_method()
	get_pointer_enetity().value_method()
}
