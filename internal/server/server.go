package server

import (
	"bytes"
	"fmt"
	"io"
	"net"

	"github.com/Israel-Andrade-P/http_from_tcp.git/internal/request"
	"github.com/Israel-Andrade-P/http_from_tcp.git/internal/response"
)

type Server struct {
	closed  bool
	handler Handler
}

type Handler func(w io.Writer, req *request.Request) *HandlerError

type HandlerError struct {
	StatusCode response.StatusCode
	Message    string
}

func runConnection(s *Server, conn io.ReadWriteCloser) {
	defer conn.Close()
	headers := response.GetDefaultHeaders(0)
	r, err := request.RequestFromReader(conn)
	if err != nil {
		response.WriteStatusLine(conn, response.StatusBadRequest)
		response.WriteHeaders(conn, headers)
		return
	}

	writer := bytes.NewBuffer([]byte{})
	handlerErr := s.handler(writer, r)
	if handlerErr != nil {
		response.WriteStatusLine(conn, handlerErr.StatusCode)
		response.WriteHeaders(conn, headers)
		conn.Write([]byte(handlerErr.Message))
		return
	}

	body := writer.Bytes()
	headers.Replace("Content-Length", fmt.Sprintf("%d", len(body)))

	response.WriteStatusLine(conn, response.StatusOK)
	response.WriteHeaders(conn, headers)
	conn.Write(body)
}

func runServer(s *Server, listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if s.closed {
			return
		}
		if err != nil {
			return
		}

		go runConnection(s, conn)
	}
}

func Serve(port uint16, handler Handler) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	server := &Server{
		closed:  false,
		handler: handler,
	}
	go runServer(server, listener)

	return server, nil
}

func (s *Server) Close() error {
	s.closed = true
	return nil
}
