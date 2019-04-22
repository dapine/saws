package main

import (
	"bufio"
	"log"
	"net"

	"github.com/dapine/saws/http/request"
	"github.com/dapine/saws/http/response"
)

const (
	httpVersion = "HTTP/1.1"
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
	header := response.NewHeader(httpVersion, req)

	_, err := req.Connection().Write([]byte(header.String()))
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
