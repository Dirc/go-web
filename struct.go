package main

import(
	"fmt"
)

type Person struct {
	fName string
	lName string
	favFood []string
}

func loop(list []string) {
	for i := range list {
		fmt.Println(list[i])
	}
}

func (person Person) walk() string {
	return fmt.Sprintln(person.fName, "is walking")
}

func main() {
	p1 := Person{"Dirc", "Grupetto", []string{"olive", "ricepie", "bananas"}}
	fmt.Println(p1)
	fmt.Println(p1.fName)
	p1.fName = "Beto"
	fmt.Println(p1)

	loop(p1.favFood)

	s := p1.walk()
	fmt.Println(s)
}