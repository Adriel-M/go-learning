package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// inits size 10 array with default 0
	var scores [10]int
	scores[2] = 339;
	fmt.Println(scores)

	// init with this array
	scores2 := [4]int{2, 3, 4, 5}
	fmt.Println(scores2)

	// pads with 0
	scores3 := [10]int{1, 2, 3}
	fmt.Println(scores3)

	// iterate
	for index, value := range scores3 {
		fmt.Println(index, value)
	}

	// slices?
	scores4 := []int{1, 2, 3, 4}
	fmt.Println(scores4)
	
	// init a slice of capacity 10 and of size 10
	scores5 := make([]int, 10)
	scores5[4] = 4
	fmt.Println(scores5)

	// init with size 0 and cap of 10
	scores6 := make([]int, 0, 10)
	// empty slice
	fmt.Println(scores6)
	// can't do this as it's out of bounds
	// scores6[4] = 4
	
	// does this create a new array?
	test1 := &scores6
	scores6 = append(scores6, 5)
	fmt.Println(scores6)
	test2 := &scores6
	// modifies the array
	fmt.Println("Is same array?", test1, test2, test1 == test2)
	// what if we want to set at 7th index?
	// breaks
	// scores6[7] = 10
	scores6 = scores6[0:8]
	scores6[7] = 10
	fmt.Println(scores6)

	// get cap for slice
	fmt.Println(cap(scores6))

	// dynamic size slices
	scores7 := make([]int, 10)
	for index := range(scores7) {
		scores7[index] = index
	}
	fmt.Println(scores7)
	currentSize := cap(scores7)
	fmt.Println("Size of scores7 is", currentSize)

	for i := 0; i < 25; i++ {
		scores7 = append(scores7, i)

		if cap(scores7) != currentSize {
			currentSize = cap(scores7)
			fmt.Println("New size is", currentSize)
		}
	}

	scores8 := make([]int, 100)
	for i := 0; i < len(scores8); i++ {
		scores8[i] = int(rand.Int31n(1000))
	}
	fmt.Println(scores8)
	sort.Ints(scores8)
	fmt.Println(scores8)

	worst := make([]int, 5)
	copy(worst, scores8[:5])
	fmt.Println(worst)

	// Maps
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]
	fmt.Println(power, exists)
	power, exists = lookup["goku"]
	fmt.Println(power, exists)

	initLookup := map[string]int {
		"goku": 9001,
		"vegeta": 2,
	}
	for key, value := range initLookup {
		fmt.Println(key, value)
	}
}

