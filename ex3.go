package main

import(
	"fmt"
)

type Person struct {
	firstname string
	lastname string
}

type secretAgent struct {
	Person
	secretAgent bool
}

func (p Person) Speak() {
	fmt.Printf("Hi, I'm %s %s\n", p.firstname, p.lastname)
}

func (sa secretAgent) Speak() {
	fmt.Printf("Hi, I'm secretAgent %s %s\n", sa.firstname, sa.lastname)
}

type humanoid interface{
	Speak()
}

func greet(h humanoid) {
	h.Speak()
}

// func vomit(h humanoid) {
// 	fmt.Printf("%T\n", h)
// 	fmt.Println(h)
// 	switch v := h.(type) {
// 	case Person:
// 		fmt.Println(v.firstname)
// 	case secretAgent:
// 		fmt.Println(v.firstname)
// 	default:
// 		fmt.Println("unknown")
// 	}
// }

func main() {
	p1 := Person{
		firstname: "Dirc",
		lastname: "Grupetto",
	}
	sa1 := secretAgent{
		Person{"James", "Bond"},
		true,
	}
	fmt.Println(p1.lastname)
	p1.Speak()
	fmt.Println(sa1.secretAgent)
	sa1.Speak()

	greet(p1)
	greet(sa1)
}