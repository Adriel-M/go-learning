package main

import (
	"fmt"
)

type Applyer func(a int, b int) int
type AddXFunc func(y int) int

type BaseAddFunc func(a int) int
type Add1Func func(a int) BaseAddFunc
type Add2Func func(a int) Add1Func

func adder(a int, b int) int {
	return a + b
}

func subber(a int, b int) int {
	return a - b
}

func applyer(a []int, f Applyer) []int {
	aLen := len(a)
	ret := make([]int, aLen)
	apply10 := func(some int) int {
		return f(some, 10)
	}
	for i := 0; i < aLen; i++ {
		ret[i] = apply10(a[i])
	}
	return ret
}

func generateAddXFunc(x int) AddXFunc {
	return func(y int) int {
		return y + x
	}
}

func curriedAdd(a int) Add2Func {
	return func(b int) Add1Func {
		return func(c int) BaseAddFunc {
			return func(d int) int {
				return a + b + c + d
			}
		}
	}
}

func main() {
	ok1 := []int{1, 2, 3}
	ok2 := make([]int, 5)
	for i := range ok2 {
		ok2[i] = i + 2
	}
	fmt.Println(applyer(ok1, adder))
	fmt.Println(applyer(ok2, subber))

	add10Func := generateAddXFunc(10)
	fmt.Println(add10Func(20))
	fmt.Println(ok2)
	ok3 := make([]int, 5)
	for i, val := range ok2 {
		ok3[i] = add10Func(val)
	}
	fmt.Println(ok3)

	step1 := curriedAdd(1)
	step2 := step1(2)
	step3 := step2(3)
	step4 := step3(4)
	fmt.Println(step1, step2, step3, step4)
}
