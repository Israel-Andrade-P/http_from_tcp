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
	parts := strings.Split(req, "\r\n")
	reqLine := strings.Split(parts[0], " ")
	if !isAllUpperAlpha(reqLine[0]) {
		return nil, fmt.Errorf("invalid method")
	}

	return nil, nil
}

func isAllUpperAlpha(str string) bool {
	for _, l := range str {
		if l < 'A' || l > 'Z' {
			return false
		}
	}
	return true
}
