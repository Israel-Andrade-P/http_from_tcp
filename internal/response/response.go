package response

import (
	"fmt"
	"io"

	"github.com/Israel-Andrade-P/http_from_tcp.git/internal/headers"
)

type Response struct{}

type StatusCode int

const (
	StatusOK            StatusCode = 200
	StatusBadRequest    StatusCode = 400
	StatusInternalError StatusCode = 500
)

func WriteStatusLine(w io.Writer, statusCode StatusCode) error {
	statusLine := []byte{}
	switch statusCode {
	case 200:
		statusLine = []byte("HTTP/1.1 200 OK\r\n")
	case 400:
		statusLine = []byte("HTTP/1.1 400 Bad Request\r\n")
	case 500:
		statusLine = []byte("HTTP/1.1 500 Internal Server Error\r\n")
	default:
		return fmt.Errorf("unrecognized status error")
	}

	_, err := w.Write(statusLine)
	if err != nil {
		return err
	}

	return nil
}

func GetDefaultHeaders(contentLen int) *headers.Headers {
	h := headers.NewHeaders()
	h.Set("Content-Length", fmt.Sprintf("%d", contentLen))
	h.Set("Connection", "close")
	h.Set("Content-Type", "text/plain")
	return h
}

func WriteHeaders(w io.Writer, h *headers.Headers) error {
	b := []byte{}
	h.ForEach(func(k, v string) {
		b = fmt.Appendf(b, "%s: %s\r\n", k, v)
	})
	b = fmt.Append(b, "\r\n")
	_, err := w.Write(b)
	return err
}
