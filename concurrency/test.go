package main

import (
	"fmt"
	"sync"
)

var (
	v  int = 0
	mu sync.Mutex
	wg sync.WaitGroup
)

func counter(value *int) {
	defer wg.Done()

	mu.Lock()
	*value = *value + 1
	fmt.Println(*value)
	mu.Unlock()
}

func main() {

	fmt.Println("start")

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go counter(&v)
	}

	wg.Wait()

	fmt.Println("final value:", v)
	fmt.Println("end")
}