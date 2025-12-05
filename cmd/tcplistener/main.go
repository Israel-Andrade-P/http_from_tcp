package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Israel-Andrade-P/http_from_tcp.git/internal/request"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalf("failed creating a listener: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("connection failed: %v", err)
		}
		fmt.Println("a new connection has been established!")

		go func(c net.Conn) {
			req, err := request.RequestFromReader(conn)
			if err != nil {
				log.Fatalf("error parsing requesr: %v", err)
			}
			fmt.Printf("- Method: %s\n", req.RequestLine.Method)
			fmt.Printf("- Target: %s\n", req.RequestLine.RequestTarget)
			fmt.Printf("- Version: %s\n", req.RequestLine.HttpVersion)
			conn.Close()
			fmt.Println("connection closed")
		}(conn)
	}
}
