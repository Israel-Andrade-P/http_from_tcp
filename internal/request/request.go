package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	req, err := parseRequestLine(string(data))
	if err != nil {
		return nil, err
	}

	return req, nil
}

func parseRequestLine(req string) (*Request, error) {
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "TRACE"}
	parts := strings.Split(req, "\r\n")
	reqLine := strings.Split(parts[0], " ")
	if len(reqLine) != 3 {
		return nil, fmt.Errorf("invalid request line")
	}
	if !isAllUpperAlpha(reqLine[0]) || !contains(methods, reqLine[0]) {
		return nil, fmt.Errorf("invalid method")
	}
	httpVersion := strings.Split(reqLine[2], "/")
	if httpVersion[1] != "1.1" {
		return nil, fmt.Errorf("invalid http version")
	}
	request := &Request{
		RequestLine: RequestLine{
			HttpVersion:   httpVersion[1],
			RequestTarget: reqLine[1],
			Method:        reqLine[0],
		},
	}

	return request, nil
}

func isAllUpperAlpha(str string) bool {
	for _, l := range str {
		if l < 'A' || l > 'Z' {
			return false
		}
	}
	return true
}

func contains(methods []string, target string) bool {
	for _, method := range methods {
		if method == target {
			return true
		}
	}
	return false
}
