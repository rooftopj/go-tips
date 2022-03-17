package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	var err error
	if i, err := func1(); i >= 0 {
		fmt.Println("in the beginning of the scope: ", err)
		// Execute other logic
		i, err = func2() // Assignment statement
		fmt.Println("in the end of the scope: ", err)
	}

	if err != nil {
		//...
	}
	fmt.Println(err)
}

func func1() (int, error) {
	i := rand.Intn(10000)
	return i, errors.New(strconv.Itoa(i))
}

func func2() (int, error) {
	i := rand.Intn(10000)
	return i, errors.New(strconv.Itoa(i))
}