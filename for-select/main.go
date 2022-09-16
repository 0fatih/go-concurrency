package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		time.Sleep(time.Second * 2)
		done <- true
	}()

loop:
	for {
		select {
		case <-done:
			fmt.Println("Heyaa!")
			break loop
		default:
			fmt.Println("Still waiting...")
		}
	}
}
