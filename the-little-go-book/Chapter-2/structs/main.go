package main

import (
	"fmt"
)

// just like C struct
type Saiyan struct {
	Name  string
	Power int
}

func modifySuper(s *Saiyan) {
	s.Power += 10000
}

type SaiyanWithFunc struct {
	Name  string
	Power int
}

func (s *SaiyanWithFunc) Super() {
	s.Power += 10000
}

func main() {
	// struct inside a func???
	type Test struct {
		Ok    int
		Value int
	}

	// Can declare function inside function
	super := func(s Saiyan) {
		s.Power += 10000
	}

	// decleration uses ,?
	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}

	fmt.Printf("%s's power level is %d\n", goku.Name, goku.Power)

	// doesn't actually modify the struct as passing in a copy
	super(goku)
	fmt.Printf("%s's power level is %d\n", goku.Name, goku.Power)

	// now it modifies
	modifySuper(&goku)
	fmt.Printf("%s's power level is %d\n", goku.Name, goku.Power)

	vegeta := SaiyanWithFunc{
		Name:  "Vegeta",
		Power: 1,
	}

	fmt.Printf("%s's power level is %d\n", vegeta.Name, vegeta.Power)
	vegeta.Super()
	fmt.Printf("%s's power level is %d\n", vegeta.Name, vegeta.Power)
}
