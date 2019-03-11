package main


import (
	"fmt"
)

func getPower() int {
	return 9001
}

func main() {
	power := getPower()
	fmt.Printf("It's over %d\n", power)
	name, level := "Goku", 9000
	fmt.Printf("%s's power level is over %d\n", name, level)
}
