package main

import (
	"fmt"
	"time"
)

func main() {
	var count int

	for i := 0; i < 100; i++ {
		go func() {

			count++
		}()
	}

	fmt.Println("Count: ", count)
	time.Sleep(1 * time.Second)
}
