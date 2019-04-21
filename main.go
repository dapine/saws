package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/dapine/saws/fs"
	"github.com/dapine/saws/request"
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

func sendResponse(req request.Request) {
	status := "HTTP/1.1 200 OK"
	t := time.Now().UTC()
	date := t.Format(time.RFC1123)
	ct := "text/html"
	content, err := fs.ReadResource(req.RequestLine().RequestUri())
	if err != nil {
		log.Println("Resource not found:", err)
	}
	cl := len(content)

	getResponseStr := fmt.Sprintf("%s\r\nDate: %s\r\nContent-Length: %d\r\nContent-Type: %s\r\n\r\n%s", status, date, cl, ct, string(content[:]))

	_, err = req.Connection().Write([]byte(getResponseStr))
	if err != nil {
		log.Println("Could not response: ", err)
	}

	req.Connection().Close()
}

func getRequest(conn net.Conn) {
	for {
		line, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			conn.Close()
			return
		}

		reqLine, err := request.RequestLineParser(string(line[:]))
		if err != nil {
			log.Println("Could not complete request: ", err)
		}

		sendResponse(request.New(conn, reqLine))
	}
}
