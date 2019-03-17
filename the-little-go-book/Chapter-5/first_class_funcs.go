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

func applyer(a []int, f Applyer) ([]int) {
	aLen := len(a)
	ret := make([]int, aLen)
	apply10 := func(some int) (int) {
		return f(some, 10)
	}
	for i := 0; i < aLen; i++ {
		ret[i] = apply10(a[i])
	}
	return ret
}

func main() {
	ok1 := []int{1, 2, 3}
	ok2 := make([]int, 5)
	for i := range(ok2) {
		ok2[i] = i + 2
	}
	fmt.Println(applyer(ok1, adder))
	fmt.Println(applyer(ok2, subber))
}