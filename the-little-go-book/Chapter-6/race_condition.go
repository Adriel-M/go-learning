package main

import (
	"fmt"
	"time"
)

var counter int = 0

func incr() {
	counter++
	fmt.Println(counter)
}

func main() {
	for i := 0; i < 200; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 100)
	// sometimes a race condition!
	fmt.Println("ended up with", counter)
}
