package main


import (
	"fmt"
)

type Applyer func(a int, b int) int

func adder(a int, b int) int {
	return a + b
}

func subber(a int, b int) int {
	return a - b
}

func applyer(a []int, f Applyer) {
	ret := make([]int, len(a))
}