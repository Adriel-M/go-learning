package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi my name is %s\n", p.Name)
}

// A saiyan is a person with a power level
type Saiyan struct {
	*Person
	Power int
}

// overloading, think js prototypes where it finds the first down the chain?
func (s *Saiyan) Introduce() {
	fmt.Printf("Hi my name is %s and my power level is %d", s.Name, s.Power)
}

type Nameless struct {
	*Person
}

func (n *Nameless) ChangeName(newName string) {
	n.Name = newName
}

type A struct {
	Value int
}

func (a *A) Test() {
	fmt.Printf("My value from A %d\n", a.Value)
}

type B struct {
	Value int
}

func (b *B) Test() {
	fmt.Printf("My value from B %d\n", b.Value)
}

type C struct {
	*A
	*B
}

func (c *C) Hi() {
	println("Hi")
}

// A constructor?
func newSaiyan(name string, power int) *Saiyan {
	return &Saiyan {
		Person: &Person {
			Name: name,
		},
		Power: power,
	}
}

func main() {
	vegeta := &Saiyan {
		Person: &Person {
			Name: "Vegeta",
		},
		Power: 9000,
	}
	vegeta.Introduce()
	fmt.Printf("My power level is %d\n", vegeta.Power)
	goku := newSaiyan("Goku", 9001)
	goku.Introduce()
	fmt.Printf("My power level is %d\n", goku.Power)

	someRandomPerson := Nameless {
		Person: &Person {},
	}
	someRandomPerson.Introduce();
	someRandomPerson.ChangeName("Bob")
	someRandomPerson.Introduce();

	someC := C{
		A: &A{1},
		B: &B{2},
	}

	someC.Hi()
	// Can't do below as Test exists in A and B
	// fmt.Println(someC.Value)
	// someC.Test()
}
