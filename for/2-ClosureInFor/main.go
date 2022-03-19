package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wrong_way()
	fmt.Println()
	correct_way()
}

func wrong_way() {
	var wg sync.WaitGroup
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8}
	wg.Add(len(tasks))
	for _, val := range tasks {
		go func() {
			// Execute the task with val
			time.Sleep(100 * time.Millisecond)
			fmt.Print(val, " ")

			wg.Done()
		}()
	}
	wg.Wait()
}

func correct_way() {
	var wg sync.WaitGroup
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8}
	wg.Add(len(tasks))
	for _, val := range tasks {
		go func(v int) {
			// Execute the task with v
			time.Sleep(100 * time.Millisecond)
			fmt.Print(v, " ")

			wg.Done()
		}(val)
	}
	wg.Wait()
}
