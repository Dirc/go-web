package main

import (
	"fmt"
	"bufio"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		// Step 2: Comment out the below
		// io.WriteString(conn, "\nHello from TCP server\n")
		// fmt.Fprintln(conn, "How is your day?")
		// fmt.Fprintf(conn, "%v", "Well, I hope!\n\n")

		// conn.Close()

		// Step 2: Instead, define a handler
		go handle(conn)
	}
}

// step 2: define handler
func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	fmt.Println("Code got here.")
}