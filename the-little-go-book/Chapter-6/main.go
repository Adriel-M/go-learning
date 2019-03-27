package main

import (
	"fmt"
	"time"
)

func process() {
	fmt.Println("Processing")
}

func main() {
	fmt.Println("start")
	// order not guaranteed with goroutines
	go process()
	go func() {
		fmt.Println("Anonymous process")
	}()
	time.Sleep(time.Millisecond * 200)
	fmt.Println("done")
}