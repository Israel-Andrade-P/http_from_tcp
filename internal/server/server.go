package server

import (
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

type Handler func(w *response.Writer, req *request.Request)

type HandlerError struct {
	StatusCode response.StatusCode
	Message    string
}

func runConnection(s *Server, conn io.ReadWriteCloser) {
	defer conn.Close()

	responseWriter := response.NewWriter(conn)
	r, err := request.RequestFromReader(conn)
	if err != nil {
		responseWriter.WriteStatusLine(response.StatusBadRequest)
		responseWriter.WriteHeaders(response.GetDefaultHeaders(len(r.Body)))
		responseWriter.WriteBody([]byte(r.Body))
		return
	}

	s.handler(responseWriter, r)
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
