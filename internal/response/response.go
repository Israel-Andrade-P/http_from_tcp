package response

import (
	"fmt"
	"io"

	"github.com/Israel-Andrade-P/http_from_tcp.git/internal/headers"
)

type Writer struct {
	writer io.Writer
}

type StatusCode int

const (
	StatusOK            StatusCode = 200
	StatusBadRequest    StatusCode = 400
	StatusInternalError StatusCode = 500
)

func NewWriter(writer io.Writer) *Writer {
	return &Writer{writer: writer}
}

func (w *Writer) WriteStatusLine(statusCode StatusCode) error {
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

	_, err := w.writer.Write(statusLine)
	if err != nil {
		return err
	}

	return nil
}

func (w *Writer) WriteHeaders(h *headers.Headers) error {
	b := []byte{}
	h.ForEach(func(k, v string) {
		b = fmt.Appendf(b, "%s: %s\r\n", k, v)
	})
	b = fmt.Append(b, "\r\n")
	_, err := w.writer.Write(b)
	return err
}

func (w *Writer) WriteBody(p []byte) (int, error) {
	n, err := w.writer.Write(p)

	return n, err
}

func GetDefaultHeaders(contentLen int) *headers.Headers {
	h := headers.NewHeaders()
	h.Set("Content-Length", fmt.Sprintf("%d", contentLen))
	h.Set("Connection", "closed")
	h.Set("Content-Type", "text/plain")
	return h
}
