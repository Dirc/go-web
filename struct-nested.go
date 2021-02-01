package main

import(
	"fmt"
)

type Vehicle struct {
	doors int
	color string
}

type Truck struct {
	vehicle Vehicle
	fourWheel bool
}

type Sedan struct {
	vehicle Vehicle
	luxury bool
}

func (truck Truck) transportationDevice() string {
  return fmt.Sprintln("Transport large objects")
}

func (sedan Sedan) transportationDevice() string {
  return fmt.Sprintln("Transport small people")
}

type transportation interface {
	transportationDevice() string
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

func main() {
	v1 := Vehicle{2, "red"}
	t1 := Truck{v1, true}

	fmt.Println(v1)
	fmt.Println(t1)
	fmt.Println(v1.color)
	fmt.Println(t1.vehicle.color)

	s1 := Sedan{Vehicle{4, "blue"}, true}
	fmt.Println(s1)
	fmt.Println(s1.vehicle.color)

	fmt.Println(t1.transportationDevice())
	fmt.Println(s1.transportationDevice())

	report(t1)
	report(s1)
}