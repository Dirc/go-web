package main

import (
	"fmt"
)

func main() {
	s := "I'm sorry dave I can't do that"
	fmt.Println(s)

	ss := []byte(s)
	sss := string(ss)
	fmt.Println(sss[0:14])
	fmt.Println(sss[10:22])
	fmt.Println(sss[17:30])

	for c := range sss {
		fmt.Println(sss[c:c+1])
	}
}