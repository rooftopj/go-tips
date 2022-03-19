package main

import "fmt"

func main() {
	print_return_val(1, func1)
	print_return_val(3, func2)
	print_return_val(5, func3)
}

func func1() (int, *int) {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	panic("Trigger error1")
	fmt.Println("panic 1")
	res1 := 1
	res2 := 2
	return res1, &res2
}

func func2() (int, *int) {
	res3 := 3
	res4 := 4
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	panic("Trigger error2")
	fmt.Println("panic 2")
	return res3, &res4
}

func func3() (res5 int, res6 *int) {
	res5 = 5
	res6 = new(int)
	*res6 = 6
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	panic("Trigger error3")
	fmt.Println("panic 3")
	return
}

func print_return_val(i int, f func() (int, *int)) {
	a, b := f()
	idx1 := i
	idx2 := i + 1
	if b == nil {
		fmt.Printf("res%d: %d, res%d: nil\n", idx1, a, idx2)
	} else {
		fmt.Printf("res%d: %d, res%d: %d\n", idx1, a, idx2, *b)
	}
	fmt.Println()
}
