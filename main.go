package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Could not listen to connections: ", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Could not accept a connection: ", err)
		}

		go getRequest(conn)
	}
}

func returnRequest(conn net.Conn) {
	getResponseStr := fmt.Sprint("HTTP/1.1 200 OK\r\nDate: Thu, 18 Apr 2019 07:47:38 GMT\r\nContent-Length:19\r\nContent-Type: text/html\r\n\r\n<h1>hello world</h1>")

	_, err := conn.Write([]byte(getResponseStr))
	if err != nil {
		log.Println("Could not response: ", err)
	}

	conn.Close()
}

func getRequest(conn net.Conn) {
	for {
		line, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			conn.Close()
			return
		}

		fmt.Println(string(line[:]))

		switch string(line[:]) {
		case "GET\r\n":
			returnRequest(conn)
			break
		case "GET / HTTP/1.1\r\n":
			returnRequest(conn)
			break
		default:
			log.Println("No implemented: ", string(line[:]))
		}
	}
}
