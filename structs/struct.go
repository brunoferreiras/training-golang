package main

import "fmt"

type Creature struct {
	Name string
	Real bool
}

func (c Creature) Dump() {
	fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}

type FlyingCreature struct {
	Creature
	WingSpan int
}

type Fooer interface {
	Foo1()
	Foo2()
	Foo3()
}

type Foo struct {

}

func (f Foo) Foo1() {
	fmt.Println("Foo1() here")
}

func (f Foo) Foo2() {
	fmt.Println("Foo2() here")
}

func (f Foo) Foo3() {
	fmt.Println("Foo3() here")
}

func main() {
	dragon := &FlyingCreature {
		Creature{"Dragon", false},
		15,
	}

	fmt.Println(dragon.Name) // Dragon
	fmt.Println(dragon.Real) // false
	fmt.Println(dragon.WingSpan) // 15
	dragon.Dump() // Name: 'Dragon', Real: false
}