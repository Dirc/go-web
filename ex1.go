package main

import(
	"fmt"
	"math"
)

type square struct {
	s float64
}

type circle struct {
  r float64
} 

func (s square) area() float64{
	return s.s * s.s
}

func (c circle) area() float64{
	pi := math.Pi
	return  pi * c.r * c.r
}

type shape interface{
  area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{4.0}
	fmt.Println(s)
	fmt.Println(s.area())

	c := circle{2.0}
	fmt.Println(c)
	fmt.Println(c.area())

	info(s)
	info(c)
}