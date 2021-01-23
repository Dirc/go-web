package main

import (
	"fmt"
)

func rot13(byteslice []byte) []byte {
	var r13 = make([]byte, len(byteslice))
	for i, v := range byteslice {
		// ascii range 97 - 122, 109 is the middle
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}

func main() {
	plaintext := "foobar"
	byteslice := []byte(plaintext)
	
	fmt.Println(byteslice)
	fmt.Println(rot13(byteslice))

}