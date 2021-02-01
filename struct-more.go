package main

import (
	"fmt"
)

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I'm a gator")
}

type flamingo bool

func (f flamingo) greeting() {
	fmt.Println("Hello, I'm pink and wonderful")
}

type swampCreature interface {
	greeting() 
}

func bayou(sc swampCreature) {
	sc.greeting()
}

func main() {
	var g1 gator 
	g1 = 1
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Printf("%T\n", x)

	x = int(g1)
	fmt.Println(x)

	g1.greeting()

	bayou(g1)

	var f flamingo
	f = true
	bayou(f)
}