package main

import (
	"fmt"
	"bufio"
	"log"
	"net"
	"strings"
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

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)

	//respond(conn)
}

// Read request method and uri
func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			httpMethod := strings.Fields(ln)[0]
			httpURI := strings.Fields(ln)[1]
			fmt.Println(httpMethod)
			fmt.Println(httpURI)

			if httpURI == "/foo" {
				foo(conn)
			}
			if httpURI == "/bar" {
				bar(conn)
			} else {
				index(conn)
			}
		}
		if ln == "" {
			break
		}
		i++

	}
}

func index(conn net.Conn) {
	body := "<!DOCTYPE html><html><head>Hello World!</head><body>Dear Grupetto</body></html>"
	httpRespond(conn, body)
}

func foo(conn net.Conn) {
	body := "<!DOCTYPE html><html><head>Hello Foo!</head><body>Foz</body></html>"
	httpRespond(conn, body)
}
func bar(conn net.Conn) {
	body := "<!DOCTYPE html><html><head>Hello Bar!</head><body>Baz</body></html>"
	httpRespond(conn, body)
}


// Respond in simple HTTP standard
func httpRespond(conn net.Conn, body string) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n") // HTTP Status
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}