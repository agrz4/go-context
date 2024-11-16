package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 1 canceled:", ctx.Err())
				return
			default:
				fmt.Println("Goroutine 1 is working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// launch another goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 2 canceled:", ctx.Err())
				return
			default:
				fmt.Println("Goroutine 2 is working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// simulate some work in the main function
	fmt.Println("Main function is working...")
	time.Sleep(2 * time.Second)

	fmt.Println("Canceling context...")
	cancel()

	// give goroutines time to finish
	time.Sleep(1 * time.Second)
	fmt.Println("Main function done.")
}
