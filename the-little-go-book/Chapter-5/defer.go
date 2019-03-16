package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("yeet")
	// exec after the func is done
	defer fmt.Println("deferred func")
	// I guess defer puts in into a stack?
	defer fmt.Println("deferred 2")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}