package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Println(err.Error())
			continue
		}

		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	defer conn.Close()
	Request(conn)
}

func Request(conn net.Conn) {
	var buffer string
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()

		if len(buffer) == 0 {
			fl := strings.Fields(ln)
			method := fl[0]
			url := fl[1]

			Response(conn, method, url)
		}

		if ln == "" {
			break
		}

		buffer = ln
	}
}

func Response(conn net.Conn, method string, url string) {
	if url == "/" {
		if method == "GET" {
			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprint(conn, "Content-Type: text/html\r\n")
			fmt.Fprint(conn, "\r\n")
			tpl.ExecuteTemplate(conn, "index.gohtml", "Index")
		}
	} else if url == "/about" {
		if method == "GET" {
			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprint(conn, "Content-Type: text/html\r\n")
			fmt.Fprint(conn, "\r\n")
			tpl.ExecuteTemplate(conn, "index.gohtml", "About")
		}
	} else if url == "/contact" {
		if method == "GET" {
			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprint(conn, "Content-Type: text/html\r\n")
			fmt.Fprint(conn, "\r\n")
			tpl.ExecuteTemplate(conn, "index.gohtml", "Contact")
		}
	} else if url == "/apply" {
		if method == "GET" {
			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprint(conn, "Content-Type: text/html\r\n")
			fmt.Fprint(conn, "\r\n")
			tpl.ExecuteTemplate(conn, "index.gohtml", "Apply")
		} else if method == "POST" {
			fmt.Fprint(conn, "HTTP/1.1 201 Created\r\n")
			fmt.Fprint(conn, "Content-Type: application/json\r\n")
			fmt.Fprint(conn, "\r\n")
			fmt.Fprint(conn, "{ \"status\": \"success\" }")
		}
	} else {
		fmt.Fprint(conn, "HTTP/1.1 404 Not Found\r\n")
	}
}
