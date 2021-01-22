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

func (p Person) pSpeak() {
	fmt.Printf("Hi, I'm %s %s\n", p.firstname, p.lastname)
}

func (sa secretAgent) saSpeak() {
	fmt.Printf("Hi, I'm secretAgent %s %s\n", sa.firstname, sa.lastname)
}

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
	p1.pSpeak()
	fmt.Println(sa1.secretAgent)
	sa1.saSpeak()
	sa1.pSpeak()
}