package main

import(
	"fmt"
)

func sliceIndex(p []int) {
	for i, j := range p {
		fmt.Println(i, j)
	}
}

func main() {
	var	primes = []int{2,3,5,7,11,13}
	fmt.Println(primes[0:2])
	sliceIndex(primes)

}