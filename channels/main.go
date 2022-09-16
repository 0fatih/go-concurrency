package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	countDown := make(chan int)

	go func() {
		defer close(countDown)
		for i := 3; i >= 0; i-- {
			countDown <- i
		}
	}()

	cont := true
	for cont {
		select {
		case val, ok := <-countDown:
			if !ok {
				cont = false
				break
			}
			fmt.Println(val)
		}
	}

	block := make(chan int)

	go func() {
		fmt.Println("Value: ", <-block)
	}()

	block <- 5

	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}

	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("\nUnblocking goroutines...")

	time.Sleep(time.Second)
	close(begin)
	wg.Wait()

	// blocking with full capacity
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	blockingStream := make(chan int, 4)
	go func() {
		defer close(blockingStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			blockingStream <- i
		}
	}()

	for integer := range blockingStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}

	clearChannelManager()
}

func clearChannelManager() {
	fmt.Println("\n\n\nClear Manager")
	chanOwner := func() <-chan int {
		resultsStream := make(chan int, 5)
		go func() {
			defer close(resultsStream)
			for i := 0; i <= 5; i++ {
				resultsStream <- i
			}
		}()
		return resultsStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving")
	fmt.Println("Clear Manager\n\n\n")
}
