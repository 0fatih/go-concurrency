package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var aritchmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		aritchmetic.Add(1)
		go func() {
			defer aritchmetic.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		aritchmetic.Add(1)
		go func() {
			defer aritchmetic.Done()
			decrement()
		}()
	}

	aritchmetic.Wait()
	fmt.Println("Arithmetic complete.")
}
