package main

import "fmt"

type Person interface {
	say()
}

type PointerImpl struct {
	words string
}

type ValueImpl struct {
	words string
}

func (p *PointerImpl) say() {
	fmt.Println("PointerImpl,", p.words)
}

func (v ValueImpl) say() {
	fmt.Println("ValueImpl,", v.words)
}

func main() {
	var pointer_PointerImpl Person = &PointerImpl{"pointer1"}

	// cannot use PointerImpl{...} (type PointerImpl) as type Person in assignment:
	// PointerImpl does not implement Person (say method has pointer receiver)
	// var value_PointerImpl Person = PointerImpl{"pointer2"}

	var pointer_ValueImpl Person = &ValueImpl{"value1"}
	var value_ValueImpl Person = ValueImpl{"value2"}

	pointer_PointerImpl.say()
	// value_PointerImpl.say()

	pointer_ValueImpl.say()
	value_ValueImpl.say()
}
