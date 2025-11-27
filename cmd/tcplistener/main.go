package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
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
			for line := range getLinesChannel(conn) {
				fmt.Printf("read: %s\n", line)
			}
			conn.Close()
			fmt.Println("connection closed")
		}(conn)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	linesChan := make(chan string)

	currentLine := ""

	go func() {
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				if err == io.EOF {
					close(linesChan)
					break
				}
				log.Fatalf("Error reading file %v", err)
			}
			//parts := strings.Split(string(buffer[:n]), "\n")
			//currentLine += strings.Join(parts[:1], " ")
			/* if len(parts) > 1 {
				linesChan <- currentLine
				currentLine = ""
				currentLine += string(parts[1])
			} */
			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				currentLine += string(data[:i])
				data = data[i+1:]
				fmt.Printf("read: %s", currentLine)
				currentLine = ""
			}
			currentLine += string(data)
		}
	}()
	return linesChan
}
