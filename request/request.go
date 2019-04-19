package request

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type requestLine struct {
	method      string
	requestUri  string
	httpVersion string
}

type Request struct {
	conn    net.Conn
	reqLine requestLine
}

const (
	GET  = "GET"
	POST = "POST"
)

// Currently supporting just HTTP 1.1
const (
	HTTP1_1 = "HTTP/1.1"
)

func New(conn net.Conn, rl requestLine) Request {
	return Request{conn: conn, reqLine: rl}
}

func (r Request) RequestLine() requestLine {
	return r.reqLine
}

func (r Request) Connection() net.Conn {
	return r.conn
}

func (rl requestLine) String() string {
	return fmt.Sprintf("%s %s %s\r\n", rl.method, rl.requestUri, rl.httpVersion)
}

func (rl requestLine) Method() string {
	return rl.method
}

func (rl requestLine) RequestUri() string {
	return rl.requestUri
}

func (rl requestLine) HttpVersion() string {
	return rl.httpVersion
}

func RequestLineParser(reqLine string) (requestLine, error) {
	reqLine = strings.Trim(reqLine, "\r\n")

	toks := strings.Split(reqLine, " ")

	if len(toks) != 3 {
		return requestLine{}, errors.New("Request line bad input")
	}

	m, err := matchRequestMethod(toks[0])
	if err != nil {
		return requestLine{}, err
	}

	// TODO: URI parser/validator
	ruri := toks[1]

	httpv, err := matchHttpVersion(toks[2])
	if err != nil {
		return requestLine{}, err
	}

	return requestLine{method: m, requestUri: ruri, httpVersion: httpv}, nil
}

func matchRequestMethod(method string) (string, error) {
	switch strings.ToUpper(method) {
	case "GET":
		return GET, nil
	case "POST":
		return POST, nil
	}

	return "", errors.New("Bad request method")
}

// Muh generics
func matchHttpVersion(version string) (string, error) {
	switch strings.ToUpper(version) {
	case "HTTP/1.1":
		return HTTP1_1, nil
	}

	return "", errors.New("Bad request HTTP version")
}
