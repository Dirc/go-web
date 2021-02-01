package main

import (
	"fmt"
)

/* func printMap(m map) {
	for key, value := range m {
		fmt.Println(key, value)
	} 
} */

func main() {
	// MAP := make(map[string]int)
	MAP := map[string]int {
		"foo": 2,
		"bar": 3,
		"baz": 5,	
	}
	MAP["dirc"] = 7
	fmt.Println(MAP)
	fmt.Println(len(MAP))

	for key := range MAP {
		fmt.Println(key)
	} 
	for value := range MAP {
		fmt.Println(MAP[value])
	} 
	for key, value := range MAP {
		fmt.Println(key, value)
	} 
}